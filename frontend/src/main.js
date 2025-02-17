import { createApp } from 'vue';
import App from './App.vue';
import { createRouter, createWebHistory } from 'vue-router';
import 'bootstrap/dist/css/bootstrap.min.css';
import 'bootstrap/dist/js/bootstrap.bundle.min.js';

// Components
import PasienList from './components/PasienList.vue';
import CreatePasien from './components/CreatePasien.vue';
import UpdatePasien from './components/UpdatePasien.vue';

const routes = [
  { path: '/', component: PasienList },
  { path: '/create', component: CreatePasien },
  { path: '/edit/:id', component: UpdatePasien },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

const app = createApp(App);
app.use(router);
app.mount('#app');
