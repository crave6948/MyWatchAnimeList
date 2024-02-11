<template>
    <div id="buttonContainer">
        <button @click="addAnime">Add Anime</button>
    </div>
    <div id="Panel">
        <input type="text" v-model="animeName" placeholder="Anime Name" />
        <input type="number" v-model="animeWatch" placeholder="Watched Episodes" />
        <input type="text" v-model="animeImage" placeholder="Anime Image" />
    </div>
</template>

<script>
const apiUrl = 'http://localhost:8080/anime';  // เปลี่ยน URL และ endpoint ตามที่คุณใช้
export default {
    data() {
        return {
            animeName: '',
            animeWatch: 0,
            animeImage: ''
        };
    },
    methods: {
        async addAnime() {
            const animeData = {
                name: this.animeName,
                watch: this.animeWatch,
                image: this.animeImage
            };

            try {
                const response = await fetch(apiUrl+'/add', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify(animeData)
                });

                if (!response.ok) {
                    throw new Error('Failed to add anime');
                }

                // Reset input fields after successful addition
                this.animeName = '';
                this.animeWatch = 0;
                this.animeImage = '';

                alert('Anime added successfully!');
            } catch (error) {
                console.error('Error adding anime:', error.message);
                alert('Failed to add anime. Please try again.');
            }
        },
    }
};
</script>

<style>
/* Add your styles here */
</style>
