<template>
  <div>
    <table>
      <thead>
        <tr>
          <th v-for="col in columns" :key="col.key">{{ col.label }}</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(row, rowIndex) in getData" :key="rowIndex">
          <td v-for="(cell, colIndex) in row" :key="colIndex">
            <template v-if="colIndex == 'image'">
              <img :src="cell" :alt="cell" style="width: 300px;">
            </template>
            <template v-else>
              {{ cell }}
            </template>
          </td>
          <td>
            <button @click="deleteAnime(row.id)">Delete</button>
            <button @click="editAnime(row)">Edit</button>
          </td>
        </tr>
      </tbody>
    </table>
    <button @click="printData">Print Data</button>
  </div>
</template>

<script>
const apiUrl = 'http://localhost:8080/anime';  // เปลี่ยน URL และ endpoint ตามที่คุณใช้
export default {
  data() {
    return {
      columns: [
        { key: 'id', label: 'ID' },
        { key: 'name', label: 'Name' },
        { key: 'episodes', label: 'Episodes' },
        { key: 'image', label: 'Image' },
      ],
      rows: [],
    };
  },
  created() {
    // ดึงข้อมูลจาก REST API
    this.fetchDataFromApi();
  },
  computed: {
    getData() {
      return this.rows;
    },
  },
  methods: {
    fetchDataFromApi() {
      fetch(apiUrl)
        .then(response => {
          if (!response.ok) {
            throw new Error(`Failed to fetch data. Status code: ${response.status}`);
          }
          return response.json();
        })
        .then(data => {
          // นำข้อมูลไปใช้ตามที่คุณต้องการ
          this.rows = data;
          console.log('Data:', data);
        })
        .catch(error => {
          console.error('Error:', error);
        });
    },
    printData() {
      console.log('Rows:', this.rows);
    },
    async deleteAnime(animeid) {
      console.log('Delete anime');
      console.log('Row:', animeid);
      const data = {
        id: animeid
      };
      try {
        const response = await fetch(apiUrl + '/remove', {
          method: 'DELETE',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify(data)
        });

        if (!response.ok) {
          throw new Error('Failed to add anime');
        }

        alert('Anime ID:' + animeid +' Remove successfully!');
      } catch (error) {
        console.error('Error While Removing anime:', error.message);
        alert('Failed to Removing anime. Please try again.');
      }
    },
    editAnime(row) {
      console.log('Edit anime');
      console.log('Row:', row);
    },
  },
};
</script>

<style scoped>
/* เพิ่มสไตล์ตามความเหมาะสม */
table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 20px;
}

th,
td {
  border: 1px solid #ddd;
  padding: 8px;
  text-align: left;
}
</style>
