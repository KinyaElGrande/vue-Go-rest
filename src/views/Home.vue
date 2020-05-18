<template>
  <div>
      <div v-for="video in videos" :key="video.ID">
        <router-link :to="{name: 'video-watch', params: {id: video.ID}}">
          <h3>{{ video.Name }}</h3>
        </router-link>
        <div v-for="Tag in video.Tags" :key="Tag.ID">
          {{ Tag.Name}}
        </div>
      </div>
  </div>
</template>

<script>
import { mapState } from 'vuex'

export default {
  name: 'Home',
  computed: {
    ...mapState(['videos'])
  }
}
</script>
