<template>
  <div :class="'guide h-full p-6 pb-20 ' + getBackground()">
    <div class="flex justify-end mr-6">
      <router-link to="/">
        <svg-icon class="mt-8 text-white" name="close" />
      </router-link>
    </div>

    <weather-display :showTemperature="false" />
    <div class="text-white mt-8">
      <h1 class="font-bold text-3xl">{{getText()}}</h1>
      <span class="inline-block mt-4 text-2xl">{{getSubtitle()}}</span>
    </div>
    <div class="mt-20 opacity-50 text-2xl negative-margin">
      <span class="font-bold">Schönwetterprogramm</span>
    </div>

    <div class="guide-slider -m-2">
      <image-slider :images="images" />
    </div>
  </div>
</template>

<script>
import WeatherDisplay from "@/components/WeatherDisplay";
import SvgIcon from "@/components/SvgIcon";
import ImageSlider from "@/components/ImageSlider";
import { mapState } from "vuex";

export default {
  name: "Guide",
  components: {
    WeatherDisplay,
    SvgIcon,
    ImageSlider
  },
  data() {
    return {
      images: getImages()
    };
  },
  computed: {
    ...mapState(["weather"])
  },
  methods: {
    getBackground() {
      if (!this.weather) {
        return;
      }
      if (this.weather.Currently.temperature < 53) {
        return "bg-cold";
      }
      switch (this.weather.Currently.icon) {
        case "sunny":
          return "bg-GornerGratGold";
        case "foggy":
          return "bg-fog";
        case "rainy":
          return "bg-rain";
        case "snow":
          return "bg-cold";
        default:
          return "bg-GornerGratGold";
      }
    },
    getText() {
      if (!this.weather) {
        return;
      }
      if (this.weather.Currently.temperature < 53) {
        return "Es wird kalt";
      }
      switch (this.weather.Currently.icon) {
        case "sunny":
          return "Das Wetter wird sonniger";
        case "foggy":
          return "Nebel zieht auf";
        case "rainy":
          return "Regen zieht auf";
        case "snow":
          return "Es wird schneien";
        default:
          return "Das Wetter wird sonniger";
      }
    },
    getSubtitle() {
      if (!this.weather) {
        return;
      }
      if (this.weather.Currently.temperature < 53) {
        return "Eine warme Jacke nicht vergessen!";
      }
      switch (this.weather.Currently.icon) {
        case "sunny":
          return "Perfekte Gelegenheit, um das Alpenpanorama zu geniessen!";
        case "foggy":
          return "Nehmen sie die Gornergrat-Bahn und entkommen sie dem Nebel.";
        case "rainy":
          return "Geniessen sie unser Alpenpanorama auch bei Regen.";
        case "snow":
          return "Geniessen sie unser Alpenpanorama im Schnee.";
        default:
          return "Perfekte Gelegenheit, um das Alpenpanorama zu geniessen!";
      }
    },
    getImages() {
      if (!this.weather) {
        return;
      }

      if (this.weather.Currently.temperature < 53) {
        return [
          {
            src: require("@/assets/cold-1.jpeg"),
            text: "Wärmen Sie sich in Ihrem Hotel auf"
          },
          {
            src: require("@/assets/cold-2.jpg"),
            text:
              "Lassen Sie sich von den Eisskulpturen am Matterhorn verzücken"
          }
        ];
      }
      switch (this.weather.Currently.icon) {
        case "sunny":
          return [
            {
              src: require("@/assets/summer-1.jpg"),
              text: "Foto vor dem Bergpanorama"
            },
            {
              src: require("@/assets/bild für jonas.png"),
              text:
                "Riffelsee, der Seelen-Spiegel des Matterhorns oberhalb von Zermatt"
            }
          ];
        case "foggy":
          return [
            {
              src: require("@/assets/fog-1.jpg"),
              text: "Das Tal verschwindet im Nebel"
            },
            {
              src: require("@/assets/fog-2.jpg"),
              text: "Geniessen Sie einen Gaumenschmauss über dem Nebel"
            }
          ];
        case "rainy":
          return [
            {
              src: require("@/assets/rainy-1.jpg"),
              text: "Besuchen Sie das grossartige Matterhorn Museum"
            },
            {
              src: require("@/assets/rainy-2.jpg"),
              text: "Lassen Sie sich verwöhnen"
            }
          ];
        case "snow":
          return [
            {
              src: require("@/assets/snow-1.jpg"),
              text: "Bestaunen Sie das schneebedeckte Zermatt"
            },
            {
              src: require("@/assets/snow-2.jpg"),
              text: "Reisen Sie zurück in ihre Kindheit"
            }
          ];
        default:
          return [
            {
              src: require("@/assets/summer-1.jpg"),
              text: "Foto vor dem Bergpanorama"
            },
            {
              src: require("@/assets/bild für jonas.png"),
              text:
                "Riffelsee, der Seelen-Spiegel des Matterhorns oberhalb von Zermatt"
            }
          ];
      }
    }
  }
};
</script>

<style lang="postcss">
.guide {
  & .negative-margin {
    margin-bottom: -1.75rem;
  }

  & .ig-image-slider-all {
    @apply mt-4;
  }

  & img.ig-image {
    min-width: inherit;
    width: 248px;
    height: 248px;
    @apply rounded-squarecard;
  }
}
</style>
