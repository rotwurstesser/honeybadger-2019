<template>
  <div>
    <svg-icon
      class="inline-block text-white weathersvgs ml-6 mt-10"
      :name="getSvgIcon"
      v-if="weather"
    />
    <span v-if="weather">{{getTemperature(weather.Currently.temperature)}}</span>
  </div>
</template>

<script>
import axios from "axios";
import SvgIcon from "@/components/SvgIcon";

export default {
  name: "WeatherDisplay",
  components: {
    SvgIcon
  },
  data() {
    return {
      weather: {}
    };
  },
  created() {
    this.weather = this.getWeather();
  },
  computed: {
    getSvgIcon() {
      if (!this.weather) {
        return;
      }
      switch (this.weather.Currently.icon) {
        case "sunny":
          return "weather-sun";
        case "foggy":
          return "weather-cloud-wind-3";
        case "rainy":
          return "weather-cloud-heavy-rain";
        default:
          return "weather-sun";
      }
    }
  },
  methods: {
    getWeather() {
      axios
        .get("http://5.9.112.47:1569/weather")
        .then(result => {
          const data = result.data;
          this.weather = data;
          console.log(this.weather);
        })
        .catch(error => {
          console.error(error);
          this.weather = {
            temperature: 10,
            icon: "sunny"
          };
        });
    },
    getTemperature() {
      // Convert Fahrenheit to celsius and round it
      return Math.round((5 / 9) * (this.weather.Currently.temperature - 32));
    }
  }
};
</script>

<style lang="postcss" scoped>
.weathersvgs {
  width: 128px;
  height: 128px;
}
</style>
