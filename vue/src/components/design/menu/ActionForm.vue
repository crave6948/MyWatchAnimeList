<template>
    <div id="buttonContainer">
        <button @click="clicked">Add Anime</button>
    </div>
    <div id="panel">
        <InputText @input-text="(e) => animeName = e" msg="Anime Name"/>
        <InputText @input-text="(e) => animeWatch = e" msg="Watched Episodes"/>
        <InputText @input-text="(e) => animeImage = e" msg="Anime Image"/>
    </div>
</template>
<script>
import InputText from '../../Misc/InputText.vue';
export default {
    props: ['sent'],
    components: {
        InputText
    },
    data() {
        return {
            animeName: '',
            animeWatch: 0,
            animeImage: ''
        }
    },
    methods: {
        clicked() {
            this.animeWatch = parseInt(this.animeWatch);
            this.sent(this.animeName, this.animeWatch, this.animeImage)
        },
    },
}   
</script>

<style scoped>
#panel {
    display: flex;
    flex-direction: column;
    align-items: center;
    margin-top: 20px;
}
</style>