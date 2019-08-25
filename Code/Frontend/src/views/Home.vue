<template>
  <div>
    <div class="home m-6">
      <div class="flex justify-between mt-8">
        <img src="@/assets/logo.png" class="inline-block" />
        <span v-if="weather" class="inline-block flex justify-between">
          <svg-icon name="icn-sun-clouds" class="inline-block" />
          <span class="text-2xl">{{getTemperature(weather.Currently.temperature) + "°C"}}</span>
        </span>
      </div>

      <div class="my-16">
        <h3 class="-mb-8 mt-12">Aktivitäten</h3>
        <router-link to="guide">
          <image-slider :images="activities" />
        </router-link>
      </div>
      <div class="mb-16">
        <h3 class="-mb-8">Angebote</h3>
        <router-link to="guide">
          <image-slider :images="discounts" />
        </router-link>
      </div>

      <div class="mb-16">
        <h3>Interaktive Map</h3>
        <iframe
          src="https://openmaptiles.github.io/positron-gl-style/#8/46.0190/7.7460"
          width="100%"
          height="480"
        ></iframe>
      </div>

      <h3 class="-mb-8">Instagram Feed</h3>
      <router-link to="guide">
        <InstagramImages hashtag="gornerbahn" />
      </router-link>
    </div>
    <svg-icon name="berg" class="w-full h-64" />
  </div>
</template>

<script>
import InstagramImages from "../components/InstagramImages.vue";
import WeatherDisplay from "../components/WeatherDisplay.vue";
import SvgIcon from "@/components/SvgIcon";
import axios from "axios";
import ImageSlider from "@/components/ImageSlider";
import { mapState } from "vuex";

export default {
  name: "home",
  components: { WeatherDisplay, InstagramImages, SvgIcon, ImageSlider },
  computed: {
    ...mapState(["weather"])
  },
  data() {
    return {
      activities: [
        {
          src: require("@/assets/riffelsee.jpg"),
          text:
            "Riffelsee, der Seelen-Spiegel des Matterhorns oberhalb von Zermatt"
        },
        {
          src: require("@/assets/sunset.jpg"),
          text: "Schaue dir den atemberaubenden Sonnenuntergang an"
        }
      ],
      discounts: [
        {
          src: require("@/assets/angebot-1.jpg"),
          text: "Registriere dich und erhalte gratis Priority Boarding"
        },
        {
          src: require("@/assets/riffelsee.jpg"),
          text:
            "Riffelsee, der Seelen-Spiegel des Matterhorns oberhalb von Zermatt"
        }
      ]
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
          this.$store.commit("setWeather", {
            temperature: 10,
            icon: "sunny"
          });
        });
    },
    getTemperature() {
      // Convert Fahrenheit to celsius and round it
      return Math.round((5 / 9) * (this.weather.Currently.temperature - 32));
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
.svg-icon--icn-sun-clouds {
  width: 40px;
  height: 35px;
}
</style>
