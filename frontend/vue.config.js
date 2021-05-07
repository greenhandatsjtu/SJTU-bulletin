module.exports = {
    "transpileDependencies": [
        "vuetify"
    ],
    devServer: {
        proxy: {
            '/': {
                target: 'http://localhost:8080',
                changeOrigin: true
            }
        }
    },
    pwa: {
        name: "SJTU Bulletin",
        themeColor: "#004098",
        msTileColor: "#ffffff",

    }
}
