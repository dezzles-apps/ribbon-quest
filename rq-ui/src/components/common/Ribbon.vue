<script setup lang="ts">
import type { PropType } from 'vue';
import { RibbonState, RibbonTheme } from '@/types/enums';

const props = defineProps({
  topText: {
    type: String
  },
  bottomText: {
    type: String
  },
  theme: {
    type: String as PropType<RibbonTheme>,
    required: true
  },
  icon: {
    type: String
  },
  achieved: {
    type: Boolean,
    default: true
  },
  highlighted: {
    type: Boolean,
    default: false
  },
  image: {
    type: String
  },
  onClick: {
    type: Function
  },
  state: {
    type: String as PropType<RibbonState>,
    default: RibbonState.Standard
  }
})

function getRibbonClass() {
  const classes: string[] = [];
  if (props.achieved) {
    classes.push('ribbon-achieved');
  } else {
    classes.push('ribbon-not-achieved');
  }
  if (props.highlighted) {
    classes.push(`ribbon-highlighted`)
  }
  classes.push(`ribbon-${props.theme.toLowerCase()}`);

  if (props.state === RibbonState.Clickable) {
    classes.push('ribbon-clickable')
  } else if (props.state === RibbonState.Loading) {
    classes.push('ribbon-loading')
  }

  return classes;
}

function isAchieved() {
  return props.achieved === undefined || props.achieved
}

function iconClasses() {
  let classes = ['icon']
  if (props.topText && props.bottomText) {
    classes.push('centre-icon')
  }
  return iconClasses
}

function onClick() {
  if (props.onClick) {
    props.onClick()
  }
}

function getStyle() {
  let s = {
    backgroundImage: ''
  }
  if (props.image) {
    s.backgroundImage = `url('${props.image}')`
  }
  return s
}
</script>

<template>
  <div style="display: inline-flex; flex-wrap: wrap;">
    <div
      class="ribbon"
      :class="getRibbonClass()"
      :style="getStyle()"
        @click="onClick()"
    >
      <div class="ribbon-top-text">
        {{ props.topText }}
      </div>
      <span :class="iconClasses()">
        <i :class="`mdi ${props.icon} mdi-36px`"></i>
      </span>
      <div :class="props.topText ? 'ribbon-bottom-text' : 'ribbon-top-text'">
        {{ props.bottomText }}
      </div>
    </div>
  </div>
</template>

<style scoped>
.ribbon {
  margin: 10px;
  position: relative;
  align-items: center;
  text-align: center;
  width: 110px;
  height: 100px;
  outline-width: 5px;
  outline-style: solid;
  border-radius: 10px;
  background-position: center;
  background-repeat: no-repeat;
  background-size: 50% auto;
}

@keyframes hilite {
    0% { }
    25% { outline-color: white; }
    50% { }
    75% { outline-color: black; }
    100% { }
    
  }
.ribbon-highlighted {
  animation: hilite 2s linear;
  animation-iteration-count: infinite;
}

.ribbon-loading {
  cursor: wait;
}

.ribbon-clickable {
  cursor: pointer;
}

.ribbon-top-text {
  font-weight: bold;
  margin-right: 10px;
  width: 100%;
  color: darkslategray;
}

.ribbon-bottom-text {
  position: absolute;
  bottom: 4px;
  left: 0;
  font-weight: bold;
  margin-right: 10px;
  width: 100%;
  color: darkslategray;
}

.icon {
  position: absolute;
}

.centre-icon {
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
}

.ribbon-not-achieved {
  opacity: 0.5;
}

.ribbon-julie {
  background-color: #DCB0F2;
  outline-color: #a35ec9;
}
.ribbon-champion {
  background-color: #66C5CC;
  outline-color: #358d95;
}
.ribbon-battle {
  background-color: #F89C74;
  outline-color: #754834;
}
.ribbon-contest {
  background-color: #FE88B1;
  outline-color: #8e3d5a;
}
.ribbon-stats {
  background-color: #9EB9F3;
  outline-color: #62749b;
}
.ribbon-shopping {
  background-color: #F6CF71;
  outline-color: #d5a32c;
}

.ribbon-loading {
  cursor: wait;
}

.ribbon-clickable {
  cursor: pointer;
}

</style>