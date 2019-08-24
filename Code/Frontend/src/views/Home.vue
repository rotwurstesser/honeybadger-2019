<template>
  <div class="home m-6">
    <WeatherDisplay :weather="weather" />
    <h3 class="-mb-8">Angebote</h3>
    <router-link to="guide">
      <InstagramImages hashtag="gornerbahn" />
    </router-link>

    <h3 class="-mb-8 mt-12">Aktivitäten</h3>
    <InstagramImages hashtag="mountainclimbing" />
    <h3 class="mt-12">Interaktive Map</h3>
    <iframe
      src="https://openmaptiles.github.io/positron-gl-style/#8/46.0190/7.7460"
      width="100%"
      height="480"
    ></iframe>
  </div>
</template>

<script>
import InstagramImages from "../components/InstagramImages.vue";
import WeatherDisplay from "../components/WeatherDisplay.vue";
import axios from "axios";
export default {
  name: "home",
  components: { WeatherDisplay, InstagramImages },
  data() {
    return {
      weather: {}
    };
  },
  created() {
    this.getWeather();
  },
  methods: {
    getWeather() {
      axios
        .get("http://5.9.112.47:1569/weather")
        .then(result => {
          const data = result.data;
          this.$store.commit("setWeather", data);
        })
        .catch(error => {
          console.error(error);
          this.weather = {
            temperature: 10,
            icon: "sunny"
          };
        });
    }
  }
};
</script>

<style lang="postcss">
.home {
  & h3 {
    @apply text-3xl font-bold;
  }
  & .weather-component {
    & span,
    & svg {
      @apply text-black;
    }
    & span {
      font-size: 35px;
      @apply ml-2;
    }
    & .weathersvgs {
      width: 60px;
      margin-top: 0;
    }
  }
  & .slider-toggler {
    color: black;
  }
}
</style>
