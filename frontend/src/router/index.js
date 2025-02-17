import { createRouter, createWebHistory } from "vue-router";
import ListPasien from "../components/PasienList.vue";
import AddPasien from "../components/AddPasien.vue";
import EditPasien from "../components/EditPasien.vue";

const routes = [
  {
    path: "/",
    name: "listPasien",
    component: ListPasien,
  },
  {
    path: "/add",
    name: "addPasien",
    component: AddPasien,
  },
  {
    path: "/edit/:id",
    name: "editPasien",
    component: EditPasien,
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
