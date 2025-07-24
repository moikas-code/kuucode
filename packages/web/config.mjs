const stage = process.env.SST_STAGE || "dev"

export default {
  url: stage === "production"
    ? "https://kuuzuki.ai"
    : `https://${stage}.kuuzuki.ai`,
  socialCard: "https://social-cards.sst.dev",
  github: "https://github.com/moikas-code/kuuzuki",
  discord: "https://kuuzuki.ai/discord",
  headerLinks: [
    { name: "Home", url: "/" },
    { name: "Docs", url: "/docs/" },
  ],
}
