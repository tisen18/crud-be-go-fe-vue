<template>
  <!-- Container utama dengan margin top 5 -->
  <div class="container mt-5">
    <!-- Judul halaman -->
    <h2>List of Pasien</h2>
    <!-- Tombol untuk menambah pasien baru -->
    <button class="btn btn-primary my-3" @click="goToCreate">Tambah Pasien</button>
    <!-- Tabel untuk menampilkan data pasien -->
    <table class="table table-bordered border-2 table-striped">
      <!-- Header tabel -->
      <thead class="text-center text-uppercase table-primary">
        <tr>
          <th>No</th>
          <th>Nama</th>
          <th>Alamat</th>
          <th>No HP</th>
          <th>Penyakit</th>
          <th>Aksi</th>
        </tr>
      </thead>
      <!-- Body tabel dengan loop data pasien -->
      <tbody>
        <tr v-for="pasien in pasienList" :key="pasien.id">
          <td>{{ pasien.id }}</td>
          <td>{{ pasien.nama }}</td>
          <td>{{ pasien.alamat }}</td>
          <td>{{ pasien.no_hp }}</td>
          <td>{{ pasien.penyakit }}</td>
          <!-- Kolom aksi dengan tombol edit dan hapus -->
          <td class="d-flex gap-3 justify-content-center">
            <button class="btn btn-warning" @click="goToEdit(pasien.id)">Edit</button>
            <button class="btn btn-danger" @click="deletePasien(pasien.id)">Hapus</button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  // Inisialisasi state komponen
  data() {
    return {
      // Array untuk menyimpan daftar pasien
      pasienList: [],
    };
  },

  // Lifecycle hook yang dipanggil saat komponen dibuat
  created() {
    // Mengambil data pasien saat komponen dibuat
    this.fetchPasien();
  },

  // Kumpulan method yang digunakan dalam komponen
  methods: {
    // Method untuk mengambil data semua pasien dari API
    fetchPasien() {
      fetch('http://localhost:9090/api/pasien')
        .then(response => response.json())
        .then(data => {
          this.pasienList = data;
        });
    },

    // Method untuk navigasi ke halaman tambah pasien
    goToCreate() {
      this.$router.push('/create');
    },

    // Method untuk navigasi ke halaman edit pasien
    goToEdit(id) {
      this.$router.push(`/edit/${id}`);
    },

    // Method untuk menghapus data pasien
    deletePasien(id) {
      fetch(`http://localhost:9090/api/pasien/${id}`, {
        method: 'DELETE',
      })
        .then(response => response.json())
        .then(data => {
          // Refresh daftar pasien setelah penghapusan
          this.fetchPasien();
        });
    },
  },
};
</script>