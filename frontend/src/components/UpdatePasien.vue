<template>
    <div class="container mt-5">
        <h2>Edit Pasien</h2>
        <form @submit.prevent="updatePasien" class="needs-validation" novalidate>
            <div class="mb-3">
                <label for="nama" class="form-label">Nama:</label>
                <input type="text" id="nama" v-model="pasien.Nama" class="form-control" required />
                <div class="invalid-feedback">Nama pasien harus diisi.</div>
            </div>

            <div class="mb-3">
                <label for="alamat" class="form-label">Alamat:</label>
                <input type="text" id="alamat" v-model="pasien.Alamat" class="form-control" required />
                <div class="invalid-feedback">Alamat pasien harus diisi.</div>
            </div>

            <div class="mb-3">
                <label for="no_hp" class="form-label">No HP:</label>
                <input type="text" id="no_hp" v-model="pasien.no_hp" class="form-control" required />
                <div class="invalid-feedback">Nomor HP pasien harus diisi.</div>
            </div>

            <div class="mb-3">
                <label for="penyakit" class="form-label">Penyakit:</label>
                <input type="text" id="penyakit" v-model="pasien.Penyakit" class="form-control" required />
                <div class="invalid-feedback">Penyakit pasien harus diisi.</div>
            </div>

            <button type="submit" class="btn btn-primary">Perbarui Pasien</button>

            <p v-if="error" class="mt-3 text-danger">{{ error }}</p>
        </form>
    </div>
</template>

<script>
export default {
    data() {
        return {
            pasien: {
                Id: '',
                Nama: '',
                Alamat: '',
                NoHP: '',
                Penyakit: '',
            },
            error: null,
        };
    },
    mounted() {
        this.fetchPasienData();
    },
    methods: {
        fetchPasienData() {
            const pasienId = this.$route.params.id;
            fetch('http://localhost:9090/api/pasien')
                .then(response => {
                    console.log('Status response:', response.status);
                    if (!response.ok) {
                        return response.text().then(text => {
                            throw new Error(`Status: ${response.status}, Message: ${text}`);
                        });
                    }
                    return response.json();
                })
                .then(data => {
                    console.log('Data yang diterima:', data);
                    const foundPasien = data.find(p => p.id === parseInt(pasienId));
                    if (!foundPasien) {
                        throw new Error('Data pasien tidak ditemukan');
                    }
                    this.pasien = {
                        Id: foundPasien.id,
                        Nama: foundPasien.nama,
                        Alamat: foundPasien.alamat,
                        no_hp: foundPasien.no_hp,
                        Penyakit: foundPasien.penyakit
                    };
                })
                .catch(error => {
                    console.error('Error lengkap:', error);
                    this.error = `Gagal mengambil data pasien: ${error.message}`;
                });
        },

        updatePasien() {
            fetch(`http://localhost:9090/api/pasien/${this.pasien.Id}`, {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    Nama: this.pasien.Nama,
                    Alamat: this.pasien.Alamat,
                    no_hp: this.pasien.no_hp,
                    Penyakit: this.pasien.Penyakit,
                }),
            })
            .then(response => {
                if (!response.ok) {
                    throw new Error('Gagal memperbarui data');
                }
                this.$router.push('/');
            })
            .catch(error => {
                console.error('Error:', error);
                this.error = error.message;
            });
        },
    },
};
</script>

<style scoped>
.error {
    color: red;
    font-weight: bold;
}
</style>