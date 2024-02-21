<template>
  <div class="anime-card">
    <img :src="imageUrl" alt="Anime Image" class="anime-card__image" />
    <h3 class="anime-card__name">{{ name }}</h3>
    <p class="anime-card__watched">Watched Episodes : {{ watched }}</p>
    <p class="anime-card__id">ID : {{ id }}</p>
    <button @click="deleteAnime">Delete</button>
    <button @click="editAnime">Edit</button>
  </div>
</template>

<script>
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
  methods: {
    deleteAnime() {
      this.calldelete(this.id)
    },
    editAnime() {
      this.calledit(this.id, this.name, this.watched, this.imageUrl)
    }
  }
}
</script>

<style scoped>
.anime-card__image {
  width: 100%;
}
</style>
