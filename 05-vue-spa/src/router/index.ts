import FeaturesPage from "@/modules/landing/pages/FeaturesPage.vue";
import HomePage from "@/modules/landing/pages/HomePage.vue";
import { createRouter, createWebHistory } from "vue-router";

export const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes:[
    {
      path: '/',
      name:'home',
      component: HomePage
    },
    {
      path: '/features',
      name:'features',
      component: FeaturesPage
    }
  ]
})

export default router

