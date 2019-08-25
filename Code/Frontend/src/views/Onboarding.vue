<template>
  <div class="onboarding p-6">
    <div class="steps mx-16 mb-6">
      <div
        @click="step = 1"
        :class="step === 1 ? 'step--active' : ''"
        class="step"
      ></div>
      <div
        @click="step = 2"
        :class="step === 2 ? 'step--active' : ''"
        class="step"
      ></div>
      <div
        @click="step = 3"
        :class="step === 3 ? 'step--active' : ''"
        class="step"
      ></div>
    </div>
    <div v-if="step === 1" class="questionWrap">
      <h2 class="text-GornerGratGold text-xl text-center font-bold mb-2">
        Was trifft zu
      </h2>
      <h3 class="text-GornerGratGold text-lg text-center">Ich bin</h3>
      <div class="choice-row">
        <div @click.prevent="setNextStep(2)" class="choice ">
          <div class="wrapping">
            <svg-icon class="text-black" name="tourist" />
            <span>Tourist</span>
          </div>
        </div>
        <div class="wrapping">
          <div @click.prevent="setNextStep(2)" class="choice">
            <svg-icon class="text-black" name="swiss" />
            <span>Schweizer</span>
          </div>
        </div>
      </div>
    </div>
    <div v-if="step === 2" class="questionWrap">
      <h2 class="text-GornerGratGold text-xl text-center font-bold mb-2">
        Mit wem bist du hier?
      </h2>
      <h3 class="text-GornerGratGold text-lg text-center">Wähle aus</h3>
      <div class="choice-row">
        <div @click.prevent="setNextStep(3)" class="choice ">
          <svg-icon class="text-black" name="trekking-person" />
          <span>Alleine</span>
        </div>
        <div @click.prevent="setNextStep(2)" class="choice">
          <svg-icon class="text-black" name="family-child" />
          <span>Mit der Familie</span>
        </div>
      </div>
      <div class="choice-row">
        <div @click.prevent="setNextStep(3)" class="choice ">
          <svg-icon class="text-black" name="family-home" />
          <span>Mit Freunden</span>
        </div>
        <div @click.prevent="setNextStep(2)" class="choice">
          <svg-icon class="text-black" name="dating-couple-balloon" />
          <span>Mit dem Partner</span>
        </div>
      </div>
    </div>

    <svg-icon name="mountainfooter1" />
  </div>
</template>

<script>
import SvgIcon from "../components/SvgIcon.vue";
export default {
  components: {
    SvgIcon
  },
  methods: {
    setNextStep(step) {
      event.currentTarget.classList.toggle("choice--active");
      // give the user some time to see the coloring/let the eyes settle
      setTimeout(() => {
        this.step = step;
      }, 100);
    }
  },
  data() {
    return {
      step: 1
    };
  }
};
</script>

<style lang="postcss">
.onboarding {
  & .wrapping {
    top: 50%;
    left: 50%;
    transform: translateY(-50%, -50%);
  }
  & .choice-row {
    @apply flex flex-row justify-around p-2 mx-2;
    & .choice {
      width: 150px;
      height: 150px;
      margin: 0 auto;

      @apply mx-1 shadow-md rounded-squarecard bg-white text-center;
      & svg {
        @apply stroke-black;
      }
    }
    & .choice--active {
      @apply bg-GornerGratGold;
      & span,
      & svg {
        @apply text-white stroke-white;
      }
    }
    & svg {
      width: 40px;
      height: 40px;
      margin: 0 auto;
    }
  }
  & .steps {
    @apply flex flex-row justify-between;
  }

  & .step {
    width: 30%;
    height: 5px;
    @apply bg-GornerGratGrey3;
    &--active {
      @apply bg-GornerGratGold;
    }
  }
  & .step--active {
    @apply bg-GornerGratGold;
  }
}
</style>
