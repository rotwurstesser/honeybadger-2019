<template>
  <div class="detail">
    <vue-video-section
      :elementId="'header-background-video'"
      :ref="'header-background-video'"
      :mp4Source="require('@/assets/offer.mp4')"
      :posterSource="require('@/assets/poster.png')"
      :mobileBreakpoint="992"
      desktopHeight="auto"
      mobileHeight="auto"
      :playsinline="true"
      :loop="true"
      :autoplay="true"
      :autobuffer="true"
      :muted="true"
    />
    <router-link to="/Guide">
      <svg-icon class="text-white absolute m-5 right-0 top-0" name="close" />
    </router-link>
    <div class="p-6">
      <h3>Wunderschöner Ausflug</h3>
      <h5 class="text-GornerGratGrey3 py-2 text-lg font-bold">
        Wunderschöner Ausflug
      </h5>
      <p>
        Lorem ipsum dolor sit amet consectetur adipisicing elit. Odio,
        reprehenderit? Lorem ipsum dolor sit amet consectetur, adipisicing elit.
      </p>
    </div>
  </div>
</template>

<script>
import SvgIcon from "../components/SvgIcon.vue";
import VueVideoSection from "vue-video-section";
export default {
  components: {
    VueVideoSection,
    SvgIcon
  },
  name: "Detail"
};
</script>

<style lang="postcss">
.detail {
  & h3 {
    @apply text-4xl font-bold;
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
