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
          <header class="col-md-4 ml-2">
            <h1 style="font-size: 4rem;" class="mt-5">WheyPal</h1>
          </header>
          <img class="col-md-6" src="../assets/weights.png" />
        </div>
        <!-- <hr class="bdr-dark-blue mb-5 mt-3"> -->
        <card>
          <div slot="header">
            <h3>Sign in</h3>
          </div>
          <ul class="list-highlight">
            <li class="my-2">
              <custom-input
                label="Email"
                placeholder="placeholder@placeholder.ca"
                v-model="email"
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
                <custom-button
                  @click="signInUser()"
                  class="mb-1"
                  color="primary"
                >
                  Sign in
                </custom-button>
                <br />
                Forgot your password?
                <br />
                <router-link to="/">
                  <custom-button class="mt-1" color="primary" outline>
                    Sucks
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
  name: "SignIn",
  components: {
    card: Card,
    "custom-input": Input,
    "custom-button": Button,
    modal: Modal
  },
  data() {
    return {
      email: "",
      password: "",
      modalIsOpen: false,
      error: ""
    };
  },
  computed: {
    ...mapState({
      userId: state => state.wheypal.userId,
      userToken: state => state.wheypal.userToken
    })
  },
  methods: {
    ...mapActions(["loginUser"]),

    async signInUser() {
      this.error = "";
      const body = {
        email: this.email,
        password: this.password
      };
      console.log(body.email);
      console.log(body.password);
      if (body.email === "" || body.password === "") {
        this.error = "Cannot sign in. Please enter all credentials.";
        this.modalIsOpen = !this.modalIsOpen;
      } else {
        try {
          await this.loginUser(body);
          this.$router.push("/discover");
        } catch (e) {
          this.error = e;
          this.modalIsOpen = !this.modalIsOpen;
        }
      }
    }
  }
};
</script>

<style scoped>
h1 {
  font-size: 3.074rem;
}
h3 {
  font-size: 2.074rem;
}
</style>
