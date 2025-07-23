export namespace Flag {
  export const KUUCODE_AUTO_SHARE = truthy("KUUCODE_AUTO_SHARE")
  export const KUUCODE_DISABLE_WATCHER = truthy("KUUCODE_DISABLE_WATCHER")

  function truthy(key: string) {
    const value = process.env[key]?.toLowerCase()
    return value === "true" || value === "1"
  }
}
