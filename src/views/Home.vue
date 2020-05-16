<template>
  <div>
      <div v-for="video in videos" :key="video.ID">
        <router-link :to="{name: 'video-watch', params: {id: video.ID}}">
          <h3>{{ video.Name }}</h3>
        </router-link>
      </div>
  </div>
</template>

<script>
import Api from '@/service/api'

export default {
  name: 'Home',
  mounted () {
    Api().get('/videos').then((response) => {
      this.videos = response.data
    })
  },
  data () {
    return {
      videos: []
    }
  }
}
</script>
