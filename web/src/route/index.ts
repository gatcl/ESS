import { createRouter, createWebHashHistory, RouteRecordRaw } from "vue-router"

const routes: RouteRecordRaw[] = [
        { path: "/", redirect: "/home" },
        {
                path: "/home",
                name: "Home",
                component: () => import("@/views/Home/Index.vue"),
        },
]

const router = createRouter({
        history: createWebHashHistory(),
        routes: routes,
})

export default router
