<template>
  <div id="app">
    <header>
      <div class="navbar navbar-dark bg-dark shadow-sm">
        <div class="container d-flex justify-content-between">
          <a href="#" class="navbar-brand d-flex align-items-center">
            <strong>Go Reddit Stream</strong>
          </a>
        </div>
      </div>
    </header>

    <main role="main">
      <div v-if="posts.length === 0">
        <section class="jumbotron text-center">
          <div class="container">
            <h1>Brace yourself</h1>
            <p class="lead text-muted">Loading your daily dose of internet...</p>
          </div>
        </section>
      </div>
      <div v-else>
        <div class="py-5 bg-light">
          <div class="container">
            <div class="row" id="row">
              <div v-for="post in posts" :key="post.imageUrl" class="col-md-4">
                <div class="card mb-4 box-shadow">
                  <img class="card-img-top" :src="post.imageUrl" alt="Card image cap">
                  <div class="card-body">
                    <p class="card-text">
                      {{post.title}}
                    </p>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<script>
import newSocket from './socket'

const MAX_POSTS_PER_PAGE = 25
export default {
  name: 'App',
  data: () => ({
    posts: []
  }),
  methods: {
    addPost(post) {
      const posts = this.posts.slice()
      if (this.posts.length > MAX_POSTS_PER_PAGE) {
        posts.pop()
      }
      posts.unshift(post)
      this.posts = posts
    }
  },
  mounted() {
    this.socket = newSocket(this.addPost)
  },
  beforeDestroy() {
    if (this.socket) {
      this.socket.close()
      this.socket = null
    }
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

.navbar {
  background: #ff4500!important;
}
</style>
