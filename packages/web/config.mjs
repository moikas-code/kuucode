const stage = process.env.SST_STAGE || "dev"

export default {
  url: stage === "production"
    ? "https://kuucode.ai"
    : `https://${stage}.kuucode.ai`,
  socialCard: "https://social-cards.sst.dev",
  github: "https://github.com/moikas-code/kuucode",
  discord: "https://kuucode.ai/discord",
  headerLinks: [
    { name: "Home", url: "/" },
    { name: "Docs", url: "/docs/" },
  ],
}
