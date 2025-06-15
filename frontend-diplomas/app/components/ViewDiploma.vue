<script setup lang="ts">
import { getAuth, onAuthStateChanged } from "firebase/auth";
import { collection, doc, setDoc, addDoc, getDoc } from "firebase/firestore";
import html2pdf from "html2pdf.js";

const diplomasStore = useDiplomasStore();
const t = useI18n();
let degreeElement = ref(null);
const nuxtApp = useNuxtApp();
const auth = nuxtApp.$auth;
const db = nuxtApp.$firestore;
const COLLECTION_NAME = "university"; // The name of your collection
const university = ref(null);

onAuthStateChanged(auth, (user) => {
  if (user) {
    user
      .getIdToken()
      .then(async (idToken) => {
        const querySnapshot = await getDoc(doc(db, COLLECTION_NAME, user.uid));
        university.value = querySnapshot.data().issuer;
      })
      .catch((error) => {
        console.error("Error getting ID token:", error);
      });
  }
});

const generatePdf = () => {
  console.log(degreeElement.value);
  if (!degreeElement.value) return;
  html2pdf(degreeElement.value, {
    margin: 10,
    filename: "degree.pdf",
    image: { type: "jpeg", quality: 0.98 },
    html2canvas: { scale: 2 },
    jsPDF: { unit: "mm", format: "a4", orientation: "portrait" },
  });
};

const viewPdf = () => {
  html2pdf()
    .from(degreeElement.value)
    .outputPdf("blob")
    .then((pdfBlob) => {
      const pdfUrl = URL.createObjectURL(pdfBlob);
      window.open(pdfUrl, "_blank"); // Open in new tab
      // Or, if you have an iframe in your template:
      // document.getElementById('pdf-viewer').src = pdfUrl;
    });
};
</script>

<template>
  <q-dialog
    v-model="diplomasStore.getViewDiplomaDialog"
    @hide="diplomasStore.setViewDiplomaDialog(false)"
  >
    <q-card style="width: 220mm; max-width: 240mm">
      <q-card-actions class="bg-primary">
        <div class="text-body1 text-white">{{ $t("diplomas.view") }}</div>
        <q-space />
        <q-btn
          flat
          round
          dense
          icon="close"
          @click="diplomasStore.setViewDiplomaDialog(false)"
          color="white"
        />
      </q-card-actions>

      <q-card-actions>
        <q-space />
        <q-btn
          unelevated
          no-caps
          color="primary"
          label="Generate PDF"
          @click="generatePdf()"
        />
        <q-btn
          unelevated
          no-caps
          color="primary"
          label="View PDF"
          @click="viewPdf()"
        />
      </q-card-actions>

      <q-card-section>
        <div class="degree-page" ref="degreeElement">
          <div class="header-section">
            <div class="university-logo-container">
              <svg
                width="50"
                height="50"
                viewBox="0 0 100 100"
                fill="none"
                xmlns="http://www.w3.org/2000/svg"
              >
                <rect x="0" y="0" width="50" height="50" fill="none" />
                <path
                  d="M50 0 L10 20 L10 70 C10 85 50 100 50 100 C50 100 90 85 90 70 L90 20 L50 0 Z"
                  fill="#0056b3"
                />
                <path d="M35 40 L65 40 L65 55 L50 60 L35 55 Z" fill="#ffffff" />
                <path
                  d="M35 40 L50 45 L65 40"
                  stroke="#ffffff"
                  stroke-width="2"
                  fill="none"
                />
                <polygon
                  points="50,25 53,30 59,30 54,34 56,40 50,37 44,40 46,34 41,30 47,30"
                  fill="#fbc531"
                />
              </svg>
            </div>
            <h1 class="university-name">
              {{ university }}
            </h1>
            <p class="university-slogan">"Knowledge, Integrity, Innovation"</p>
          </div>

          <div class="main-content">
            <p class="declaration-text">
              By the authority of the Board of Trustees, and on the
              recommendation of the Faculty,
            </p>
            <p class="declaration-text">
              this institution proudly confers upon
            </p>

            <h2 class="student-name">{{ diplomasStore.getDiploma.holder.name }}</h2>

            <p class="declaration-text">the degree of</p>
            <h3 class="degree-title">
              {{ diplomasStore.getDiploma.degree_name }}
            </h3>
            <p class="field-of-study">
              in {{ diplomasStore.getDiploma.holder }}
            </p>

            <div class="honor-section">
              <p class="honor-title">
                with all the rights, honors, and privileges appertaining
                thereto,
              </p>
              <p class="honor-text">
                graduated with Distinction:
                {{ diplomasStore.getDiploma.grade_point }}
              </p>
            </div>
          </div>

          <div class="seal-section"></div>
          <div class="footer-section">
            <div class="signature-block">
              <div class="signature-line president"></div>
              <p class="official-title">University President</p>
              <p class="official-name">Dr. Eleanor Vance</p>
            </div>
            <div class="signature-block">
              <div class="signature-line dean"></div>
              <p class="official-title">Dean of the College of Engineering</p>
              <p class="official-name">Prof. Robert Chen</p>
            </div>
          </div>

          <div class="degree-details">
            <div class="degree-id">
              <p
                class="text-caption text-grey-8 q-mb-xs"
                style="font-size: 12px"
              >
                Degree ID: {{ diplomasStore.getDiploma.hash }}
              </p>
              <p
                class="text-caption text-grey-8 q-mb-xs"
                style="font-size: 12px"
              >
                Issuer: {{ university }}
              </p>
            </div>
            <div class="issue-date">
              <p
                class="text-caption text-grey-8 q-mb-xs"
                style="font-size: 12px"
              >
                Issued: {{ diplomasStore.getDiploma.issued_date }}
              </p>
            </div>
          </div>
        </div>
      </q-card-section>
    </q-card>
  </q-dialog>
</template>

<style scoped>
.degree-page {
  width: 25 0mm; /* A4 width */
  min-height: 275mm; /* A4 height */
  background-color: #fff;
  box-shadow: 0 10px 25px rgba(0, 0, 0, 0.2);
  position: relative;
  padding: 30px;
  box-sizing: border-box;
}


.header-section {
  text-align: center;
  margin-bottom: 30px;
}

.university-logo {
  max-width: 180px;
  height: auto;
  margin-bottom: 15px;
}

.university-name {
  font-family: "Playfair Display", serif;
  font-size: 2.8em;
  color: #003366; /* Darker blue */
  margin: 0;
  line-height: 1.1;
}

.university-slogan {
  font-size: 1.1em;
  color: #555;
  margin-top: 5px;
  font-style: italic;
}

.main-content {
  text-align: center;
  margin-top: 30px;
}

.declaration-text {
  font-size: 1.3em;
  color: #333;
  margin-bottom: 25px;
  line-height: 1.6;
}

.student-name {
  font-family: "Playfair Display", serif;
  font-size: 3.5em;
  color: #0056b3;
  margin: 20px 0;
  font-weight: 700;
  text-transform: capitalize;
}

.degree-title {
  font-family: "Playfair Display", serif;
  font-size: 2.2em;
  color: #003366;
  margin-top: 15px;
  margin-bottom: 10px;
}

.field-of-study {
  font-size: 1.6em;
  color: #444;
  margin-bottom: 25px;
  font-weight: bold;
}

.honor-section {
  margin-top: 30px;
  font-size: 1.1em;
  color: #666;
}

.honor-title {
  font-family: "Playfair Display", serif;
  font-size: 1.5em;
  color: #0056b3;
  margin-bottom: 10px;
}

.seal-section {
  position: absolute;
  bottom: 60px;
  left: 50%;
  transform: translateX(-50%);
  width: 120px; /* Adjust size as needed */
  height: 120px;
  background-image: url("university_seal.png"); /* Placeholder for your seal image */
  background-size: contain;
  background-repeat: no-repeat;
  background-position: center;
  opacity: 0.2; /* Watermark effect */
  z-index: 0; /* Ensure it's behind text */
}

.footer-section {
  display: flex;
  justify-content: space-around;
  align-items: flex-end;
  margin-top: 50px;
  position: relative;
  z-index: 1; /* Ensure text is above seal */
}

.signature-block {
  flex: 1;
  text-align: center;
  margin: 0 15px;
  padding-top: 20px;
  border-top: 1px solid #ccc;
}

.signature-line {
  height: 50px; /* Space for signature image */
  margin-bottom: 5px;
  background-size: contain;
  background-repeat: no-repeat;
  background-position: center;
}

.signature-line.president {
  background-image: url("president_signature.png");
}

.signature-line.dean {
  background-image: url("dean_signature.png");
}

.official-title {
  font-size: 1em;
  font-weight: bold;
  color: #333;
  margin-bottom: 5px;
}

.official-name {
  font-size: 0.9em;
  color: #666;
}

.degree-details {
  margin-top: 40px;
  text-align: left;
  font-size: 0.95em;
  color: #555;
  display: flex;
  justify-content: space-between;
  align-items: flex-end;
}

.degree-id {
  flex-grow: 1;
}

.issue-date {
  text-align: right;
  flex-grow: 1;
}
</style>
