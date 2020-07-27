<template>
    <v-lazy v-model="visible"
            min-height="100"
            transition="scroll-y-reverse-transition">
        <v-card
                light
                :href="notice['href']"
                raised
                hover
                target="_blank"
                :color="title(notice['_type'])[1]"
                class="mx-0 my-4">
            <v-card-title v-text="notice['title']"></v-card-title>
            <v-row class="ml-1" dense>
                <v-col>
                    <v-chip class="mb-1"
                            label
                            small
                            color="grey darken-1"
                            text-color="white"
                    >
                        <v-icon left>mdi-label</v-icon>
                        {{title(notice['_type'])[0]}}
                    </v-chip>
                </v-col>
                <v-col class="mr-2 text-right">
                    {{notice['date'].slice(0,10)}}
                </v-col>
            </v-row>
        </v-card>
    </v-lazy>
</template>

<script>
    export default {
        name: "NoticeItem",
        props: ['notice'],
        data: () => ({
            visible: false,
        }),
        methods: {
            title: function (_type) {
                // 生成可读的类型
                switch (_type) {
                    case 'jwc':
                        return ['教务处', 'blue lighten-3']
                    case 'xsb':
                        return ['学生办', 'teal lighten-3']
                    case 'partTime':
                        return ['实习信息', 'green lighten-3']
                    case 'fullTime':
                        return ['全职招聘', 'deep-orange lighten-3']
                }
            }
        }
    }
</script>

<style scoped>

</style>
