<template>
  <div id="container">
    <Box v-for="anime in data" :key="anime.id">
      <AnimeCard
        :imageUrl="anime.image"
        :name="anime.name"
        :watched="anime.watch"
        :id="anime.id"
        :calldelete="deleteAnime"
      />
    </Box>
  </div>
  <ActionMenu :addAnime="addAnime" />
</template>

<script>
import Box from './Misc/Box.vue'
import AnimeCard from './design/AnimeCard.vue'
import ActionMenu from './ActionMenu.vue'

const apiUrl = 'http://localhost:8080/anime' // เปลี่ยน URL และ endpoint ตามที่คุณใช้

export default {
  name: 'AnimeTableNew',
  components: {
    Box,
    AnimeCard,
    ActionMenu
  },
  data() {
    return {
      data: []
    }
  },
  created() {
    // ดึงข้อมูลจาก REST API
    this.fetchDataFromApi()
  },
  computed: {
    getData() {
      return this.data
    }
  },
  methods: {
    fetchDataFromApi() {
      fetch(apiUrl)
        .then((response) => {
          if (!response.ok) {
            throw new Error(`Failed to fetch data. Status code: ${response.status}`)
          }
          return response.json()
        })
        .then((data) => {
          // นำข้อมูลไปใช้ตามที่คุณต้องการ
          this.data = data
          console.log('Data:', data)
        })
        .catch((error) => {
          console.error('Error:', error)
        })
    },
    printData() {
      console.log('Rows:', this.data)
    },
    async deleteAnime(animeid) {
      console.log('Delete anime')
      console.log('Row:', animeid)
      const data = {
        id: animeid
      }
      try {
        const response = await fetch(apiUrl + '/remove', {
          method: 'DELETE',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify(data)
        })

        if (!response.ok) {
          throw new Error('Failed to add anime')
        }

        alert('Anime ID:' + animeid + ' Remove successfully!')
        this.fetchDataFromApi()
      } catch (error) {
        console.error('Error While Removing anime:', error.message)
        alert('Failed to Removing anime. Please try again.')
      }
    },
    editAnime(row) {
      console.log('Edit anime')
      console.log('Row:', row)
    },
    async addAnime(animeName, animeWatch, animeImage) {
      const animeData = {
        name: animeName,
        watch: animeWatch,
        image: animeImage
      }

      try {
        const response = await fetch(apiUrl + '/add', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify(animeData)
        })

        if (!response.ok) {
          throw new Error('Failed to add anime')
        }

        // Reset input fields after successful addition
        this.animeName = ''
        this.animeWatch = 0
        this.animeImage = ''

        alert('Anime added successfully!')
        this.fetchDataFromApi()
      } catch (error) {
        console.error('Error adding anime:', error.message)
        alert('Failed to add anime. Please try again.')
      }
    }
  }
}
</script>

<style scoped>
#container {
  display: grid;
  grid-template-columns: repeat(auto-fit, 400px);
  justify-content: center;
}
/* Your component styles go here */
</style>
