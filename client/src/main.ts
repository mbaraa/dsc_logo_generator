import { createApp } from "vue";
import App from "./App.vue";
import store from "./store/store";

export const backendAddress = "https://logogen.sheev.xyz"

createApp(App)
    .use(store)
    .mount('#app')
