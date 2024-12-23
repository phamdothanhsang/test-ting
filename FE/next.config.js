module.exports = {
    reactStrictMode: true,
    async rewrites() {
      return [
        {
          source: '/kols/:path*',
          destination: 'http://localhost:8081/kols/:path*',
        },
      ]
    }
  }
  