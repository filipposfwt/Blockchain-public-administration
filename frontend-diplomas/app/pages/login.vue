<script setup lang="ts">
import { useFirebaseAuth } from '#imports';

const { signIn, user, authReady } = useFirebaseAuth();

definePageMeta({
  layout: "auth",
  name: "login",
  middleware: ['auth']
});

const t = useI18n();
const router = useRouter();
const email = ref("");
const password = ref("");
const loading = ref(false);
const error = ref("");

const handleLogin = () => {
  signIn(email.value, password.value).then((res) => {
    if (res.accessToken) {
      localStorage.setItem("dms_accessToken", res.accessToken);
      router.push({ name: "diplomas" });
    }
  });
};


</script>

<template>  
  <q-page class="flex flex-center">
    <q-card
      square
      class="shadow-0 text-black"
      style="width: 400px; max-width: 80vw"
    >
    <q-card-section class="text-left bg-primary">
        <div class="text-body1 text-white">{{$t('title')}}</div>
      </q-card-section>
      <q-card-section class="text-left">
        <p class="text-h6  ">{{ $t("auth.login") }}</p>
        <q-form @submit="handleLogin">
          <p class="q-mb-xs text-weight-bold">{{ $t("auth.email") }}</p>
          <q-input
            v-model="email"
            :placeholder="$t('auth.email')"
            type="email"
            class="q-mb-md"
            :rules="[(val) => !!val || 'Email is required']"
            filled
            square
            dense
            bg-color="grey-2"
          />
          <p class="q-mb-xs text-weight-bold">{{ $t("auth.password") }}</p>
          <q-input
            v-model="password"
            :placeholder="$t('auth.password')"
            type="password"
            :rules="[(val) => !!val || 'Password is required']"
            dense
            square
            bg-color="grey-2"
          />

          <div v-if="error" class="text-negative text-caption q-mt-sm">
            {{ error }}
          </div>

          <q-btn
            :label="$t('auth.login')"
            type="submit"
            color="primary"
            :loading="loading"
            unelevated
            no-caps
            class="full-width q-mt-lg"
          />
        </q-form>
      </q-card-section>

   
    </q-card>
  </q-page>
</template>

<style scoped>
.logo {
  filter: brightness(0) saturate(100%) invert(12%) sepia(25%) saturate(2994%)
    hue-rotate(225deg) brightness(95%) contrast(111%);
}
</style>
