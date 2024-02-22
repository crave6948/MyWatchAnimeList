<template>
  <div class="anime-card">
    <div v-if="editmode">
      <div style="display: flex; justify-content: center; align-items: center; flex-direction: column;">
        <InputText @input-text="(e) => (editedname = e)" :msg="editedname" :initvalue="editedname"/>
        <InputText @input-text="(e) => (editedwatched = convertToInt(e))" :msg="editedwatched.toString()" :initvalue="editedwatched.toString()"/>
        <InputText @input-text="(e) => (editedimageUrl = e)" msg="Image Url" :initvalue="editedimageUrl"/>
        <div>
          <button @click="saveAnime">Save</button>
          <button @click="editmode = false">Cancel</button>
        </div>
      </div>
    </div>
    <div v-else>
      <img :src="imageUrl" alt="Anime Image" class="anime-card__image" />
      <h3 class="anime-card__name">{{ name }}</h3>
      <p class="anime-card__watched">Watched Episodes : {{ watched }}</p>
      <p class="anime-card__id">ID : {{ id }}</p>
      <div style="display: flex; justify-content: center;">
        <button @click="deleteAnime">Delete</button>
        <button @click="editAnime">Edit</button>
      </div>
    </div>
  </div>
</template>

<script>
import InputText from '../Misc/InputText.vue'
export default {
  props: {
    imageUrl: {
      type: String,
      required: true
    },
    name: {
      type: String,
      required: true
    },
    watched: {
      type: Number,
      required: true
    },
    id: {
      type: String,
      required: true
    },
    calldelete: {
      type: Function,
      required: true
    },
    calledit: {
      type: Function,
      required: true
    }
  },
  data() {
    return { 
      editmode: false,
      editedname : this.name,
      editedwatched : this.watched,
      editedimageUrl : this.imageUrl
    }
  },
  methods: {
    deleteAnime() {
      this.calldelete(this.id)
    },
    editAnime() {
      this.editmode = !this.editmode
    },
    saveAnime() {
      if (!this.checkVaild()) {
        alert('Invalid Input')
        return
      }
      if (!this.checkNumber()) {
        alert('Watched Episodes must be a number')
        return
      }
      this.calledit(this.id, this.editedname, this.editedwatched, this.editedimageUrl)
      this.editmode = false
    },
    checkVaild() {
      if (this.editedname === '' || this.editedwatched === 0 || this.editedimageUrl === '') {
        return false
      }
      return true
    },
    checkNumber() {
      if (isNaN(this.editedwatched) || this.editedwatched < 0 || typeof this.editedwatched !== 'number') {
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
    }
  },
  components: {
    InputText
  }
}
</script>

<style scoped>
.anime-card__image {
  width: 100%;
}
</style>
