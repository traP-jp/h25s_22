import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
    },
    {
      path: '/about',
      name: 'about',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../views/AboutView.vue'),
    },
    // Rooms edit routes
    {
      path: '/rooms/edit/date-and-time',
      name: 'rooms-edit-date-and-time',
      component: () => import('../views/rooms/edit/DateAndTimeView.vue'),
    },
    {
      path: '/rooms/edit/place',
      name: 'rooms-edit-place',
      component: () => import('../views/rooms/edit/PlaceView.vue'),
    },
    {
      path: '/rooms/edit/decide',
      name: 'rooms-edit-decide',
      component: () => import('../views/rooms/edit/DecideView.vue'),
    },
    {
      path: '/rooms/edit/created',
      name: 'rooms-edit-created',
      component: () => import('../views/rooms/edit/CreatedView.vue'),
    },
    // Rooms with ID routes
    {
      path: '/rooms/:room_id/places',
      name: 'rooms-places',
      component: () => import('../views/rooms/PlacesView.vue'),
    },
    {
      path: '/rooms/:room_id/choose-time',
      name: 'rooms-choose-time',
      component: () => import('../views/rooms/ChooseTimeView.vue'),
    },
    {
      path: '/rooms/:room_id/reorder-by-preference',
      name: 'rooms-reorder-by-preference',
      component: () => import('../views/rooms/ReorderByPreferenceView.vue'),
    },
    {
      path: '/rooms/:room_id/vote-results',
      name: 'rooms-vote-results',
      component: () => import('../views/rooms/VoteResultsView.vue'),
    },
    {
      path: '/rooms/:room_id/room-participation',
      name: 'rooms-participate',
      component: () => import('../views/rooms/RoomParticipation.vue'),
    },
  ],
})

export default router
