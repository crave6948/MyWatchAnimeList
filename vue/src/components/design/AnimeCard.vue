<template>
  <div class="anime-card">
    <img :src="imageUrl" alt="Anime Image" class="anime-card__image" />
    <h3 class="anime-card__name">{{ name }}</h3>
    <p class="anime-card__watched">Watched Episodes : {{ watched }}</p>
    <p class="anime-card__id">ID : {{ id }}</p>
    <button @click="deleteAnime">Delete</button>
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
    }
  },
  methods: {
    deleteAnime() {
      this.calldelete(this.id)
    }
  }
}
</script>

<style scoped>
.anime-card__image {
  width: 100%;
}
</style>
