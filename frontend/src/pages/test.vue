<script setup lang="ts">
import { ref, computed } from 'vue'
import { 
  X, FileText, Clock, CheckCircle, Menu, 
  Bell, Search
} from 'lucide-vue-next'

// --- TYPES ---
type Status = 'pending' | 'approved' | 'rejected'

interface Document {
  id: string
  title: string
  type: 'ใบขออนุมัติ' | 'ใบลา' | 'โครงการ'
  submitter: {
    name: string
    avatar: string
  }
  submitDate: string
  status: Status
  reason?: string
  reviewDate?: string
  amount?: number
}

// --- STATE ---
const brand = {
  primary: '#8B001D', // Dark Red
  light: '#A52A2A'    // Lighter Red
}

const activeTab = ref<'pending' | 'history'>('pending')
const selectedDoc = ref<Document | null>(null)
const showModal = ref(false)
const modalType = ref<'approve' | 'reject' | 'edit'>('approve')
const reason = ref('')
const sidebarOpen = ref(true)
const searchTerm = ref('')

const user = {
  name: 'คุณผู้จัดการ',
  avatar: `https://ui-avatars.com/api/?name=Manager&background=random&color=fff`
}

const documents = ref<Document[]>([
  { id: 'DOC001', title: 'ใบขออนุมัติซื้ออุปกรณ์สำนักงาน', type: 'ใบขออนุมัติ', submitter: { name: 'สมชาย ใจดี', avatar: `https://ui-avatars.com/api/?name=สมชาย+ใจดี&background=random&color=fff` }, submitDate: '2025-11-08', status: 'pending', amount: 15000 },
  { id: 'DOC002', title: 'ใบลาพักร้อน 5 วัน', type: 'ใบลา', submitter: { name: 'สมหญิง รักษ์ดี', avatar: `https://ui-avatars.com/api/?name=สมหญิง+รักษ์ดี&background=random&color=fff` }, submitDate: '2025-11-09', status: 'pending' },
  { id: 'DOC003', title: 'ขออนุมัติงบประมาณโครงการ ABC', type: 'โครงการ', submitter: { name: 'วิชัย สุขใจ', avatar: `https://ui-avatars.com/api/?name=วิชัย+สุขใจ&background=random&color=fff` }, submitDate: '2025-11-10', status: 'pending', amount: 250000 },
  { id: 'DOC004', title: 'ใบขออนุมัติเดินทางต่างจังหวัด', type: 'ใบขออนุมัติ', submitter: { name: 'ประยุทธ มั่นคง', avatar: `https://ui-avatars.com/api/?name=ประยุทธ+มั่นคง&background=random&color=fff` }, submitDate: '2025-11-05', status: 'approved', reviewDate: '2025-11-06', reason: 'อนุมัติตามที่เสนอ', amount: 5000 },
  { id: 'DOC005', title: 'ใบลาป่วย 3 วัน', type: 'ใบลา', submitter: { name: 'มานี ระวัง', avatar: `https://ui-avatars.com/api/?name=มานี+ระวัง&background=random&color=fff` }, submitDate: '2025-11-03', status: 'rejected', reviewDate: '2025-11-04', reason: 'ไม่มีใบรับรองแพทย์ประกอบ' },
])

// --- COMPUTED ---
const filteredDocs = computed(() => {
  const tabFiltered = documents.value.filter(d => {
    if (activeTab.value === 'pending') return d.status === 'pending'
    if (activeTab.value === 'history') return d.status !== 'pending'
    return true
  })
  if (!searchTerm.value) {
    return tabFiltered
  }
  return tabFiltered.filter(d => 
    d.title.toLowerCase().includes(searchTerm.value.toLowerCase()) ||
    d.submitter.name.toLowerCase().includes(searchTerm.value.toLowerCase())
  )
})

const pendingDocs = computed(() => documents.value.filter(d => d.status === 'pending'))

// --- METHODS ---
function openModal(doc: Document, type: 'approve' | 'reject') {
  selectedDoc.value = doc
  modalType.value = type
  reason.value = doc.reason || ''
  showModal.value = true
}

function getStatusInfo(status: Status) {
  switch (status) {
    case 'pending': return { label: 'รอพิจารณา', class: 'bg-warning/10 text-warning-content border-warning/20', icon: Clock }
    case 'approved': return { label: 'อนุมัติ', class: 'bg-success/10 text-success-content border-success/20', icon: CheckCircle }
    case 'rejected': return { label: 'ปฏิเสธ', class: 'bg-error/10 text-error-content border-error/20', icon: X }
  }
}

function getDocTypeIcon(type: Document['type']) {
  switch (type) {
    case 'ใบขออนุมัติ': return FileText
    case 'ใบลา': return Clock
    case 'โครงการ': return FileText
  }
}

function formatCurrency(amount?: number) {
  if (amount === undefined) return ''
  return new Intl.NumberFormat('th-TH', { style: 'currency', currency: 'THB' }).format(amount)
}

function submitAction() {
  if (!selectedDoc.value) return
  
  const docIndex = documents.value.findIndex(d => d.id === selectedDoc.value!.id)
  if (docIndex === -1) return

  const newStatus = modalType.value === 'approve' ? 'approved' : 'rejected'
  
  const docToUpdate = documents.value[docIndex];
  if (!docToUpdate) return;

  docToUpdate.status = newStatus;
  docToUpdate.reason = reason.value || (newStatus === 'approved' ? 'อนุมัติ' : undefined);
  docToUpdate.reviewDate = new Date().toISOString().split('T')[0];
  
  showModal.value = false
}
</script>

<template>
  <div class="flex h-screen bg-base-200 font-sans">
    <!-- Sidebar -->
    <aside 
      class="flex flex-col bg-base-100 text-base-content transition-all duration-300 ease-in-out"
      :class="sidebarOpen ? 'w-72' : 'w-20'"
    >
      <div class="flex items-center justify-between p-4 h-20 border-b border-base-300">
        <div v-if="sidebarOpen" class="flex items-center gap-3">
          <div class="p-2 rounded-lg" :style="{ backgroundColor: brand.primary }">
            <FileText class="w-6 h-6 text-white" />
          </div>
          <span class="font-bold text-xl">อนุมัติเอกสาร</span>
        </div>
        <button @click="sidebarOpen = !sidebarOpen" class="btn btn-ghost btn-circle">
          <Menu class="w-6 h-6" />
        </button>
      </div>

      <nav class="flex-1 p-4 space-y-2">
        <button
          @click="activeTab = 'pending'"
          class="w-full flex items-center gap-3 px-4 py-3 rounded-lg font-semibold transition-all duration-300"
          :class="activeTab === 'pending' ? `bg-primary text-primary-content shadow-md` : 'hover:bg-base-200'"
        >
          <Clock class="w-5 h-5" />
          <span v-if="sidebarOpen" class="flex-1 text-left">รอพิจารณา</span>
          <span v-if="sidebarOpen && pendingDocs.length > 0" class="badge badge-secondary">{{ pendingDocs.length }}</span>
        </button>
        <button
          @click="activeTab = 'history'"
          class="w-full flex items-center gap-3 px-4 py-3 rounded-lg font-semibold transition-all duration-300"
          :class="activeTab === 'history' ? `bg-primary text-primary-content shadow-md` : 'hover:bg-base-200'"
        >
          <CheckCircle class="w-5 h-5" />
          <span v-if="sidebarOpen" class="flex-1 text-left">ประวัติ</span>
        </button>
      </nav>

      <div class="p-4 border-t border-base-300">
        <div v-if="sidebarOpen" class="flex items-center gap-3">
          <div class="avatar">
            <div class="w-10 rounded-full">
              <img :src="user.avatar" />
            </div>
          </div>
          <div>
            <p class="font-semibold">{{ user.name }}</p>
            <a href="#" class="text-xs text-error/70 hover:text-error font-semibold">ออกจากระบบ</a>
          </div>
        </div>
      </div>
    </aside>

    <!-- Main Content -->
    <div class="flex-1 flex flex-col overflow-hidden">
      <!-- Header -->
      <header class="flex items-center justify-between p-4 h-20 bg-base-100 border-b border-base-300">
        <div class="flex items-center">
          <h1 class="text-2xl font-bold text-base-content">{{ activeTab === 'pending' ? 'เอกสารรอพิจารณา' : 'ประวัติการพิจารณา' }}</h1>
        </div>
        <div class="flex items-center gap-4">
          <div class="form-control">
            <label class="input input-bordered flex items-center gap-2 rounded-full h-12">
              <Search class="w-5 h-5 text-base-content/50" />
              <input type="text" class="grow" placeholder="ค้นหา..." v-model="searchTerm" />
            </label>
          </div>
          <button class="btn btn-ghost btn-circle">
            <div class="indicator">
              <Bell class="w-6 h-6" />
              <span class="badge badge-xs badge-primary indicator-item"></span>
            </div>
          </button>
        </div>
      </header>

      <!-- Content -->
      <main class="flex-1 overflow-y-auto p-6 md:p-8">
        <Transition name="fade" mode="out-in">
          <div :key="activeTab">
            <div v-if="filteredDocs.length === 0" class="flex flex-col items-center justify-center h-full text-center text-base-content/60 pt-16">
              <div class="w-48 h-48 mb-6">
                <svg viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" class="opacity-20">
                  <path d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
                </svg>
              </div>
              <h3 class="text-2xl font-bold mb-2">
                {{ searchTerm ? 'ไม่พบเอกสาร' : (activeTab === 'pending' ? 'ไม่มีเอกสารรอพิจารณา' : 'ไม่มีประวัติ') }}
              </h3>
              <p>{{ searchTerm ? 'ลองค้นหาด้วยคำอื่น' : 'ทุกอย่างเรียบร้อยดี!' }}</p>
            </div>

            <div v-else class="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 2xl:grid-cols-4 gap-6">
              <div
                v-for="doc in filteredDocs"
                :key="doc.id"
                class="card bg-base-100 shadow-md hover:shadow-xl transition-all duration-300 ease-in-out transform hover:-translate-y-1"
              >
                <div class="card-body p-5">
                  <div class="flex justify-between items-start mb-3">
                    <div class="flex items-center gap-3">
                      <div class="avatar">
                        <div class="w-10 h-10 rounded-full">
                          <img :src="doc.submitter.avatar" />
                        </div>
                      </div>
                      <div>
                        <p class="font-semibold text-base-content">{{ doc.submitter.name }}</p>
                        <p class="text-xs text-base-content/60">{{ doc.submitDate }}</p>
                      </div>
                    </div>
                    <div class="badge badge-outline text-xs" :class="getStatusInfo(doc.status).class">
                      {{ getStatusInfo(doc.status).label }}
                    </div>
                  </div>

                  <h2 class="card-title text-base mb-2 truncate">{{ doc.title }}</h2>
                  
                  <div class="flex items-center justify-between text-sm text-base-content/80 mb-4">
                    <div class="flex items-center gap-2">
                      <component :is="getDocTypeIcon(doc.type)" class="w-4 h-4 opacity-60" />
                      <span>{{ doc.type }}</span>
                    </div>
                    <div v-if="doc.amount" class="font-bold text-primary">
                      {{ formatCurrency(doc.amount) }}
                    </div>
                  </div>

                  <div v-if="activeTab === 'pending'" class="card-actions justify-end border-t border-base-300 pt-4 mt-auto">
                    <button class="btn btn-sm btn-ghost text-error" @click="openModal(doc, 'reject')">ปฏิเสธ</button>
                    <button class="btn btn-sm btn-primary" @click="openModal(doc, 'approve')">อนุมัติ</button>
                  </div>
                  
                  <div v-if="activeTab === 'history' && doc.reason" class="text-xs p-3 bg-base-200 rounded-lg mt-2">
                    <p><span class="font-semibold">เหตุผล:</span> {{ doc.reason }}</p>
                    <p class="text-base-content/60">วันที่พิจารณา: {{ doc.reviewDate }}</p>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </Transition>
      </main>
    </div>

    <!-- Modal -->
    <dialog class="modal" :class="{ 'modal-open': showModal }">
      <div class="modal-box">
        <button @click="showModal = false" class="btn btn-sm btn-circle btn-ghost absolute right-2 top-2">✕</button>
        <h3 class="font-bold text-lg mb-4">
          {{ modalType === 'approve' ? 'ยืนยันการอนุมัติ' : 'ยืนยันการปฏิเสธ' }}
        </h3>
        
        <div v-if="selectedDoc" class="space-y-2 mb-4 p-4 bg-base-200 rounded-box">
          <p><strong>เอกสาร:</strong> {{ selectedDoc.title }}</p>
          <p><strong>ผู้ส่ง:</strong> {{ selectedDoc.submitter.name }}</p>
        </div>

        <div class="form-control">
          <label class="label">
            <span class="label-text">เหตุผล (ไม่บังคับ)</span>
          </label>
          <textarea v-model="reason" class="textarea textarea-bordered" rows="3" placeholder="ระบุเหตุผล..."></textarea>
        </div>
        
        <div class="modal-action">
          <button class="btn" @click="showModal = false">ยกเลิก</button>
          <button 
            class="btn" 
            :class="modalType === 'approve' ? 'btn-success' : 'btn-error'"
            @click="submitAction"
          >
            ยืนยัน
          </button>
        </div>
      </div>
      <div class="modal-backdrop" @click="showModal = false"></div>
    </dialog>
  </div>
</template>

<style src="./test.css"></style>
