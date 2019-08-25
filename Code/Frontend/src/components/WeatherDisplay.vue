<template>
  <div class="weather-component">
    <svg-icon class="inline-block text-white weathersvgs mt-8" :name="getSvgIcon" v-if="weather" />
    <span v-if="weather && showTemperature">
      {{
      getTemperature(weather.Currently.temperature) + "°C"
      }}
    </span>
  </div>
</template>

<script>
import SvgIcon from "@/components/SvgIcon";
import { mapState } from "vuex";

export default {
  name: "WeatherDisplay",
  components: {
    SvgIcon
  },
  props: {
    showTemperature: {
      required: false,
      default: true
    }
  },
  computed: {
    ...mapState(["weather"]),
    getSvgIcon() {
      if (!this.weather && this.weather.Currently) {
        return;
      }
      if (this.weather.Currently.temperature < 53 && !this.showTemperature) {
        return "temperature-thermometer-down";
      }
      switch (this.weather.Currently.icon) {
        case "sunny":
          return "weather-sun";
        case "foggy":
          return "weather-cloud-wind-3";
        case "rainy":
          return "weather-cloud-heavy-rain";
        case "snow":
          return "ice-snowflake-invert";
        default:
          return "weather-sun";
      }
    }
  },
  methods: {
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
