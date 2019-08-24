<template>
  <div></div>
</template>

<script>
import axios from "axios";
export default {
  name: "WeatherDisplay",
  data() {
    return {
      weather: "sunny"
    };
  },
  created() {
    this.weather =
      localStorage.getItem("weather") || this.getWeather() || "sunny";
  },
  methods: {
    getWeather() {
      axios
        .get(
          "https://api.darksky.net/forecast/5375e67e1f87323c4b9b9ef18c29f5db/46.020714,7.749117"
        )
        .then(result => {
          console.log(result);

          debugger;
        });
    }
  }
};
</script>

<style lang="scss" scoped></style>
