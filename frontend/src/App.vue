<template>
    <v-app>
        <v-app-bar
                app
                color="indigo"
                dark
        >
            <div class="d-flex align-center">
                <v-img
                        alt="SJTU Logo"
                        class="shrink mr-2"
                        contain
                        :src="require('@/assets/logo.png')"
                        transition="scale-transition"
                        width="40"
                />
                <v-toolbar-title class="font-italic font-weight-bold">SJTU Bulletin</v-toolbar-title>
            </div>

            <v-spacer></v-spacer>

            <v-btn
                    href="https://github.com/greenhandatsjtu/SJTU-bulletin"
                    target="_blank"
                    text
            >
                <span class="mr-2">GitHub</span>
                <v-icon>mdi-open-in-new</v-icon>
            </v-btn>
        </v-app-bar>

        <v-main>
            <Bulletin/>
        </v-main>
    </v-app>
</template>

<script>
    import Bulletin from "./views/Bulletin";

    export default {
        name: 'App',

        components: {
            Bulletin
        },

        data: () => ({
            //
        }),
    };
</script>
