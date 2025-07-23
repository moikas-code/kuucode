// Prevents additional console window on Windows in release, DO NOT REMOVE!!
#![cfg_attr(not(debug_assertions), windows_subsystem = "windows")]

use tauri::{Manager, Emitter};
use std::process::{Command, Stdio, Child};
use std::sync::{Arc, Mutex};
use std::io::{BufRead, BufReader};
use std::thread;

#[derive(Clone, serde::Serialize)]
struct Payload {
    args: Vec<String>,
    cwd: String,
}

// Global state to track the kuucode process
struct AppState {
    kuucode_process: Arc<Mutex<Option<Child>>>,
}

#[tauri::command]
async fn start_kuucode_tui(
    state: tauri::State<'_, AppState>,
    window: tauri::Window,
) -> Result<String, String> {
    let mut process_guard = state.kuucode_process.lock().unwrap();
    
    // Kill existing process if running
    if let Some(mut child) = process_guard.take() {
        let _ = child.kill();
    }
    
    // Start kuucode TUI
    let mut child = Command::new("kuucode")
        .stdin(Stdio::piped())
        .stdout(Stdio::piped())
        .stderr(Stdio::piped())
        .spawn()
        .map_err(|e| format!("Failed to start kuucode: {}. Make sure kuucode is installed.", e))?;

    // Get stdout and stderr handles
    let stdout = child.stdout.take().unwrap();
    let stderr = child.stderr.take().unwrap();
    
    // Clone window for the threads
    let window_stdout = window.clone();
    let window_stderr = window.clone();
    
    // Spawn thread to read stdout
    thread::spawn(move || {
        let reader = BufReader::new(stdout);
        for line in reader.lines() {
            if let Ok(line) = line {
                let _ = window_stdout.emit("kuucode-output", line);
            }
        }
    });
    
    // Spawn thread to read stderr
    thread::spawn(move || {
        let reader = BufReader::new(stderr);
        for line in reader.lines() {
            if let Ok(line) = line {
                let _ = window_stderr.emit("kuucode-error", line);
            }
        }
    });
    
    // Store the process
    *process_guard = Some(child);
    
    Ok("Kuucode TUI started successfully".to_string())
}

#[tauri::command]
async fn send_input_to_kuucode(
    state: tauri::State<'_, AppState>,
    input: String,
) -> Result<(), String> {
    let process_guard = state.kuucode_process.lock().unwrap();
    
    if let Some(child) = process_guard.as_ref() {
        if let Some(stdin) = child.stdin.as_ref() {
            use std::io::Write;
            let mut stdin = stdin;
            stdin.write_all(input.as_bytes())
                .map_err(|e| format!("Failed to send input: {}", e))?;
            stdin.flush()
                .map_err(|e| format!("Failed to flush input: {}", e))?;
        }
    }
    
    Ok(())
}

#[tauri::command]
async fn stop_kuucode_tui(state: tauri::State<'_, AppState>) -> Result<String, String> {
    let mut process_guard = state.kuucode_process.lock().unwrap();
    
    if let Some(mut child) = process_guard.take() {
        child.kill()
            .map_err(|e| format!("Failed to stop kuucode: {}", e))?;
        Ok("Kuucode stopped".to_string())
    } else {
        Ok("Kuucode was not running".to_string())
    }
}

// Simplified - no menu for now to get basic functionality working

fn main() {
    tauri::Builder::default()
        .manage(AppState {
            kuucode_process: Arc::new(Mutex::new(None)),
        })
        .invoke_handler(tauri::generate_handler![
            start_kuucode_tui,
            send_input_to_kuucode,
            stop_kuucode_tui
        ])
        .setup(|app| {
            let window = app.get_webview_window("main").unwrap();
            
            // Set window properties
            window.set_title("Kuucode").unwrap();
            
            // Get current directory
            let current_dir = std::env::current_dir()
                .unwrap_or_default()
                .to_string_lossy()
                .to_string();
            
            // Emit ready event
            window.emit("app-ready", Payload {
                args: std::env::args().collect(),
                cwd: current_dir,
            }).unwrap();
            
            Ok(())
        })
        // Menu events removed for now
        .on_window_event(|window, event| {
            match event {
                tauri::WindowEvent::CloseRequested { .. } => {
                    // Clean up kuucode process when window closes
                    if let Some(state) = window.try_state::<AppState>() {
                        let mut process_guard = state.kuucode_process.lock().unwrap();
                        if let Some(mut child) = process_guard.take() {
                            let _ = child.kill();
                        }
                    }
                }
                _ => {}
            }
        })
        .run(tauri::generate_context!())
        .expect("error while running tauri application");
}