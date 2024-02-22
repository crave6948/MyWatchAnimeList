<template>
  <div id="buttonContainer">
    <button @click="clicked">Add Anime</button>
  </div>
  <div id="panel">
    <InputText @input-text="(e) => (animeName = e)" msg="Anime Name" />
    <InputText @input-text="(e) => (animeWatch = convertToInt(e))" msg="Watched Episodes" :initvalue="animeWatch.toString()"/>
    <InputText @input-text="(e) => (animeImage = e)" msg="Anime Image" />
  </div>
</template>
<script>
import InputText from '../../Misc/InputText.vue'
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
    checkVaild() {
      if (this.animeName === '' || this.animeWatch === 0 || this.animeImage === '') {
        return false
      }
      return true
    },
    checkNumber() {
      if (isNaN(this.animeWatch) || this.animeWatch < 0 || typeof this.animeWatch !== 'number') {
        return false
      }
      return true
    },
    convertToInt(value) {
      try {
        return parseInt(value)
      } catch (error) {
        console.error(error)
      }
      return 0
    },
    clicked() {
      if (!this.checkVaild()) {
        alert('Please fill in all the fields')
        return
      }
      if (!this.checkNumber()) {
        alert('Watched Episodes must be a number')
        return
      }
      try {
        this.animeWatch = this.convertToInt(this.animeWatch)
      } catch (error) {
        console.error(error)
      }
      this.sent(this.animeName, this.animeWatch, this.animeImage)
    }
  }
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
