<template>
  <div class="container">
    <div class="row">
      <div class="col-md-12">
        <modal
          id="basicModal"
          :active="modalIsOpen"
          title="Error"
          :show-button="false"
          @toggle="modalIsOpen = false"
        >
          <p>{{ this.error }}</p>
        </modal>
        <div class="row">
          <header class="col-md-4">
            <h1 class="mt-5 ml-2">WheyPal</h1>
          </header>
          <img class="col-md-7" src="../assets/cyclists.png" />
        </div>
        <!-- <hr class="bdr-dark-blue mb-5 mt-3"> -->
        <card>
          <div slot="header">
            <h3>Sign up</h3>
          </div>
          <ul class="list-highlight">
            <li class="my-2">
              <custom-input
                v-model="name"
                label="Name"
                placeholder="Name"
              ></custom-input>
            </li>
            <li class="my-2">
              <custom-input
                v-model="email"
                label="Email"
                placeholder="placeholder@placeholder.ca"
              ></custom-input>
            </li>
            <li class="my-2">
              <custom-input
                v-model="password"
                label="Password"
                placeholder="Password"
              ></custom-input>
            </li>
          </ul>
          <div slot="footer">
            <div class="row ml-1">
              <div class="col-md-4">
                <custom-button @click="signUp()" class="mb-1" color="primary">
                  Sign up
                </custom-button>
                <br />
                Already a member?
                <br />
                <router-link to="/signIn">
                  <custom-button class="mt-1" color="primary" outline>
                    Sign in
                  </custom-button>
                </router-link>
              </div>
            </div>
          </div>
        </card>
      </div>
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
import {
  Card,
  Input,
  Button,
  Modal
} from "rbc-wm-framework-vuejs/dist/wm/components";
import { mapActions, mapState } from "vuex";
export default {
  name: "Home",
  components: {
    card: Card,
    "custom-input": Input,
    "custom-button": Button,
    modal: Modal
  },
  data() {
    return {
      name: "",
      email: "",
      password: "",
      modalIsOpen: false,
      error: "",
      image: "cyclists.png"
    };
  },
  computed: {
    ...mapState({
      userId: state => state.wheypal.userId
    })
  },
  methods: {
    ...mapActions(["createUser", "logoffUser"]),

    async signUp() {
      this.error = "";
      const body = {
        name: this.name,
        email: this.email,
        password: this.password
      };
      if (body.name === "" || body.email === "" || body.password === "") {
        this.error = "Cannot sign up. Please enter all credentials.";
        this.modalIsOpen = !this.modalIsOpen;
      } else {
        try {
          await this.createUser(body);
          this.$router.push("/discover");
        } catch (e) {
          this.error = e;
          this.modalIsOpen = !this.modalIsOpen;
        }
      }
    },
    mounted() {
      this.logoffUser(); // clear all states
    }
  }
};
</script>

<style scoped>
h1 {
  font-size: 4rem;
}
h3 {
  font-size: 2.074rem;
}
.header {
  position: relative;
  text-align: center;
}
</style>
