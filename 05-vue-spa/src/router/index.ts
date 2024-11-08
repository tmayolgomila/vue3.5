import LoginPage from "@/modules/auth/pages/LoginPage.vue";
import LandingLayout from "@/modules/landing/layouts/LandingLayout.vue";
import ContactPage from "@/modules/landing/pages/ContactPage.vue";
import FeaturesPage from "@/modules/landing/pages/FeaturesPage.vue";
import HomePage from "@/modules/landing/pages/HomePage.vue";
import PricingPage from "@/modules/landing/pages/PricingPage.vue";
import { createRouter, createWebHistory } from "vue-router";

export const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes:[
    {
      path: "/",
      name: 'landing',
      component: LandingLayout,
      children:[
        {
          path:'/',
          name:'home',
          component: HomePage
        },
        {
          path:'/features',
          name:'features',
          component: FeaturesPage
        },
        {
          path:'/pricing',
          name:'pricing',
          component: PricingPage
        },
        {
          path:'/contact',
          name:'contact',
          component: ContactPage
        },
      ]
    },
    {
      path: '/auth',
      name:'auth',
      component: LoginPage
    }
  ]
})

export default router




