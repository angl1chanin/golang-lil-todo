<template>
  <div class="todo-item">
    <div class="todo-item__body">
      <input
          class="todo-item__inp-cbx"
          :id="item.title+item.id"
          :checked="item.completed"
          @click="this.$emit('updateStatus', item.id)"
          type="checkbox">
      <label
          :for="item.title+item.id"
      >
        <p class="todo-item__body-text">{{ item.title }}</p>
      </label>
    </div>
    <div class="todo-item__manage">
      <input class="todo-item__grouped-row-cbx" type="checkbox" :id="item.id" style="display: none;">
      <ul class="todo-item__grouped-row">
        <label
            class="todo-item__swipe-btn"
            :for="item.id"
        >
          <svg viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <g id="SVGRepo_bgCarrier" stroke-width="0"></g>
            <g id="SVGRepo_tracerCarrier" stroke-linecap="round" stroke-linejoin="round"></g>
            <g id="SVGRepo_iconCarrier">
              <path d="M15 7L10 12L15 17" stroke="#000000" stroke-width="1.5"
                    stroke-linecap="round" stroke-linejoin="round"></path>
            </g>
          </svg>
        </label>
        <li class="todo-item__grouped-row-item">
          <button
              class="todo-item__grouped-row-item-btn todo-item__grouped-row-item-btn_delete"
              @click="this.$emit('deleteTask', item.id)"
          >
            <svg width="22" height="24" viewBox="0 0 22 24" fill="nonee"
                 xmlns="http://www.w3.org/2000/svg">
              <path
                  d="M7.66602 20.7432C7.84863 20.7432 7.99528 20.6906 8.10596 20.5854C8.21663 20.4803 8.26921 20.342 8.26367 20.1704L7.90674 8.19238C7.90674 8.02083 7.84863 7.88525 7.73242 7.78564C7.62174 7.68604 7.4834 7.63623 7.31738 7.63623C7.12923 7.63623 6.97982 7.6888 6.86914 7.79395C6.764 7.89355 6.71419 8.02913 6.71973 8.20068L7.06006 20.1787C7.07113 20.3503 7.12923 20.4858 7.23438 20.5854C7.34505 20.6906 7.48893 20.7432 7.66602 20.7432ZM10.9365 20.7432C11.1191 20.7432 11.2658 20.6906 11.3765 20.5854C11.4927 20.4803 11.5508 20.342 11.5508 20.1704V8.20068C11.5508 8.02913 11.4927 7.89355 11.3765 7.79395C11.2658 7.6888 11.1191 7.63623 10.9365 7.63623C10.7539 7.63623 10.6045 7.6888 10.4883 7.79395C10.3776 7.89355 10.3223 8.02913 10.3223 8.20068V20.1704C10.3223 20.342 10.3776 20.4803 10.4883 20.5854C10.6045 20.6906 10.7539 20.7432 10.9365 20.7432ZM14.1987 20.7432C14.3813 20.7432 14.5252 20.6934 14.6304 20.5938C14.741 20.4941 14.7992 20.3558 14.8047 20.1787L15.145 8.20898C15.1506 8.03743 15.098 7.89909 14.9873 7.79395C14.8822 7.6888 14.7383 7.63623 14.5557 7.63623C14.3841 7.63623 14.243 7.68604 14.1323 7.78564C14.0216 7.88525 13.9635 8.0236 13.958 8.20068L13.6094 20.1787C13.6038 20.3447 13.6536 20.4803 13.7588 20.5854C13.8695 20.6906 14.0161 20.7432 14.1987 20.7432ZM6.31299 4.99658V2.63086C6.31299 1.88932 6.53711 1.30827 6.98535 0.887695C7.43913 0.461589 8.06169 0.248535 8.85303 0.248535H12.9951C13.7865 0.248535 14.4062 0.461589 14.8545 0.887695C15.3083 1.30827 15.5352 1.88932 15.5352 2.63086V4.99658H14.0825V2.71387C14.0825 2.38184 13.9774 2.11621 13.7671 1.91699C13.5568 1.71777 13.2718 1.61816 12.9121 1.61816H8.93604C8.57633 1.61816 8.29134 1.71777 8.08105 1.91699C7.87077 2.11621 7.76562 2.38184 7.76562 2.71387V4.99658H6.31299ZM1.33252 5.71875C1.1499 5.71875 0.989421 5.65234 0.851074 5.51953C0.712728 5.38118 0.643555 5.21794 0.643555 5.02979C0.643555 4.84163 0.712728 4.67839 0.851074 4.54004C0.989421 4.40169 1.1499 4.33252 1.33252 4.33252H20.5322C20.7204 4.33252 20.8809 4.40169 21.0137 4.54004C21.152 4.67285 21.2212 4.8361 21.2212 5.02979C21.2212 5.21794 21.1548 5.38118 21.022 5.51953C20.8892 5.65234 20.7259 5.71875 20.5322 5.71875H1.33252ZM6.06396 23.9556C5.57699 23.9556 5.14811 23.856 4.77734 23.6567C4.41211 23.4575 4.12158 23.1781 3.90576 22.8184C3.68994 22.4587 3.57096 22.0409 3.54883 21.5649L2.76025 5.50293H19.1128L18.3242 21.5566C18.3021 22.0326 18.1803 22.4504 17.959 22.8101C17.7432 23.1698 17.4499 23.4492 17.0791 23.6484C16.7139 23.8532 16.2933 23.9556 15.8174 23.9556H6.06396Z" />
            </svg>
          </button>
          <button
              class="todo-item__grouped-row-item-btn todo-item__grouped-row-item-btn_edit"
          >
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 48 48" width="48px">
              <path d="M 36 5.0097656 C 34.205301 5.0097656 32.410791 5.6901377 31.050781 7.0507812 L 8.9160156 29.183594 C 8.4960384 29.603571 8.1884588 30.12585 8.0253906 30.699219 L 5.0585938 41.087891 A 1.50015 1.50015 0 0 0 6.9121094 42.941406 L 17.302734 39.974609 A 1.50015 1.50015 0 0 0 17.304688 39.972656 C 17.874212 39.808939 18.39521 39.50518 18.816406 39.083984 L 40.949219 16.949219 C 43.670344 14.228094 43.670344 9.7719064 40.949219 7.0507812 C 39.589209 5.6901377 37.794699 5.0097656 36 5.0097656 z M 36 7.9921875 C 37.020801 7.9921875 38.040182 8.3855186 38.826172 9.171875 A 1.50015 1.50015 0 0 0 38.828125 9.171875 C 40.403 10.74675 40.403 13.25325 38.828125 14.828125 L 36.888672 16.767578 L 31.232422 11.111328 L 33.171875 9.171875 C 33.957865 8.3855186 34.979199 7.9921875 36 7.9921875 z M 29.111328 13.232422 L 34.767578 18.888672 L 16.693359 36.962891 C 16.634729 37.021121 16.560472 37.065723 16.476562 37.089844 L 8.6835938 39.316406 L 10.910156 31.521484 A 1.50015 1.50015 0 0 0 10.910156 31.519531 C 10.933086 31.438901 10.975086 31.366709 11.037109 31.304688 L 29.111328 13.232422 z"/>
            </svg>
          </button>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
export default {
  name: "TodoItem",
  props: {
    item: {
      type: Object,
      required: true,
    }
  },
  methods: {

  }
}
</script>

<style lang="scss" scoped>
/* todo-item */
.todo-item {
  $primary: #03a333;
  $gray: #b9b8c3;

  display: flex;
  align-items: center;
  min-height: 30px;
  overflow: hidden;

  &__body {
    margin-right: auto;
    width: 100%;
    display: flex;

    & > * {
      cursor: pointer;
    }
  }

  &__body label {
    display: inline-block;
    width: 100%;
  }

  &__body-text {
    color: var(--text-color);
    font-size: 14px;
  }

  &__inp-cbx {
    margin-right: 12px;
  }

  &__inp-cbx:checked ~ label p {
    color: #b9b8c3;
  }

  &__manage {
    position: relative;
  }

  &__swipe-btn {
    cursor: pointer;
    text-align: center;
    width: 25px;
    height: 25px;
    border-radius: 50%;
    display: flex;
    align-items: center;
    justify-content: center;
    transition: background-color .2s linear;
    margin-right: 3px;

    &:hover {
      background-color: #dddddd;
    }

    svg {
      width: 20px;
      height: 20px;
      transform: scale(0.9);
      transition: transform 0.2s linear;
    }
  }

  &__grouped-row {
    display: flex;
    position: absolute;
    top: 50%;
    transform: translateY(-50%);
    right: -95px;
    transition: right 0.3s linear;
  }

  // &__grouped-row-item {
  // }
  &__grouped-row-cbx:checked ~ &__grouped-row &__swipe-btn {
    background-color: #dddddd;
  }

  &__grouped-row-cbx:checked ~ &__grouped-row {
    right: 2px;
  }

  &__grouped-row-cbx:checked ~ &__grouped-row &__swipe-btn svg {
    transform: rotate(180deg);
  }

  &__grouped-row-item {
    display: flex;
    gap: 3px;
  }

  &__grouped-row-item-btn {
    background-color: none;
    border: none;
    outline: none;
    cursor: pointer;
    padding: 1px 15px;
    border-radius: 3px;
    display: flex;
    align-items: center;
    justify-content: center;
    transition: background-color .2s linear;
    svg {
      width: 16px;
    }

    svg path {
      fill: white;
    }

    &_delete {
      background-color: #FF453A;
      &:hover {
        background-color: #ff271b;
      }
    }

    &_edit {
      background-color: #AEAEB2;
      &:hover {
        background-color: #89898d;
      }
    }
  }
}
</style>