import 'material-design-icons-iconfont/dist/material-design-icons.css' // Ensure you are using css-loader
import Vue from 'vue'
import Vuetify from 'vuetify';
import 'vuetify/dist/vuetify.min.css';

Vue.use(Vuetify);

export default new Vuetify({
    theme: {
        options: {
            customProperties: true,
        },
        themes: {
            light: {
                primary: '#0277BD',
                secondary: '#01579B',
                accent: '#1DE9B6',
                error: '#FF5252',
                success: '#4CAF50',
                warning: '#FFC107'
            },
        },
        customProperties: true,
        iconfont: 'md',
    },
});