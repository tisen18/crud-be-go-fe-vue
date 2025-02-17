<template>
    <!-- Container utama dengan margin top 5 -->
    <div class="container mt-5">
      <!-- Judul form -->
      <h2>Create Pasien</h2>
      <!-- Form untuk membuat pasien baru -->
      <form class="form-group" @submit.prevent="createPasien">
        <!-- Input field untuk nama -->
        <div class="mb-3">
          <label class="form-label" for="nama">Nama : </label>
          <input class="form-control" type="text" id="nama" v-model="pasien.Nama" required />
        </div>
        <!-- Input field untuk alamat -->
        <div class="mb-3">
          <label class="form-label" for="alamat">Alamat:</label>
          <input class="form-control" type="text" id="alamat" v-model="pasien.Alamat" required />
        </div>
        <!-- Input field untuk nomor HP -->
        <div class="mb-3">
          <label class="form-label" for="no_hp">No HP:</label>
          <input class="form-control" type="text" id="no_hp" v-model="pasien.no_hp" required />
        </div>
        <!-- Input field untuk penyakit -->
        <div class="mb-3">
          <label class="form-label" for="penyakit">Penyakit:</label>
          <input class="form-control" type="text" id="penyakit" v-model="pasien.Penyakit" required />
        </div>
        <!-- Tombol submit -->
        <button class="btn btn-primary" type="submit">Simpan Pasien</button>
      </form>
    </div>
  </template>
  
  <script>
  export default {
    // Inisialisasi state komponen
    data() {
      return {
        // Object untuk menyimpan data pasien baru
        pasien: {
          Nama: '',
          Alamat: '',
          NoHP: '',
          Penyakit: '',
        },
        // State untuk menampung pesan error
        error: null,
      };
    },
  
    methods: {
      // Method untuk membuat pasien baru
      async createPasien() {
        try {
          // Melakukan request POST ke API
          const response = await fetch('http://localhost:9090/api/pasien/', {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json',
            },
            body: JSON.stringify(this.pasien),
          });
  
          // Mengecek apakah request berhasil
          if (!response.ok) {
            const errorText = await response.text();
            throw new Error(`Error: ${response.statusText}, ${errorText}`);
          }
  
          // Mengambil data response
          const data = await response.json();
          // Redirect ke halaman utama setelah berhasil
          this.$router.push('/');
        } catch (error) {
          // Menangani error
          this.error = error.message;
          console.error('Error:', error);
        }
      },
    },
  };
  </script>