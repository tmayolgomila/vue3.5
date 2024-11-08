import AuthLayout from "@/modules/auth/layouts/AuthLayout.vue";
import LoginPage from "@/modules/auth/pages/LoginPage.vue";
import RegisterPage from "@/modules/auth/pages/RegisterPage.vue";
import NotFound404 from "@/modules/common/pages/NotFound404.vue";
import LandingLayout from "@/modules/landing/layouts/LandingLayout.vue";
import ContactPage from "@/modules/landing/pages/ContactPage.vue";
import FeaturesPage from "@/modules/landing/pages/FeaturesPage.vue";
import HomePage from "@/modules/landing/pages/HomePage.vue";
import PricingPage from "@/modules/landing/pages/PricingPage.vue";
import PokemonPage from "@/modules/pokemons/pages/PokemonPage.vue";
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
        {
          path: '/pkm/:id',
          name:'pkm',
          props: (route)=>{
            const id = +route.params.id
            console.log({id})
            return isNaN(id) ? {id: 1} : { id }
          },
          component: PokemonPage
        }
      ]
    },
    {
      path: '/auth',
      redirect: {name: 'login'},
      component: AuthLayout,
      children:[
        {
          path: 'login',
          name: 'login',
          component:LoginPage
        },
        {
          path: 'register',
          name: 'register',
          component: RegisterPage
        }
      ]
    },
    {
      path: '/:pathMatch(.*)*',
      component: NotFound404
    },
  ]
})

export default router




