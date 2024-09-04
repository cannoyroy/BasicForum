import { createRouter, createWebHistory } from 'vue-router'
import LoginView from '../views/LoginView.vue'
import AboutView from '../views/RegView.vue'
import StudentBoard from '../views/ROOT/StudentBoard.vue'
import StudentPost from '../views/Student/StudentPost.vue'
import StudentGetPost from '../views/Student/StudentGetPost.vue'
import StudentPutPost from '../views/Student/StudentPutPost.vue'
import StudentDelePost from '../views/Student/StudentDelePost.vue'
import StudentReportPost from '../views/Student/StudentReportPost.vue'
import StudentGetReport from '../views/Student/StudentGetReport.vue'
import TeacherBoard from '../views/ROOT/TeacherBoard.vue'
import AdminGetReport from '../views/Teacher/AdminGetReport.vue'
import AdminProReport from '../views/Teacher/AdminProReport.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/login',
      name: 'logon',
      component: LoginView
    },
    {
      path: '/reg',
      name: 'reg',
      component: AboutView
    },
    {
      path: '/board/student',
      name: 'StudentBoard',
      component: StudentBoard,
      children: [
        {
          path: 'post',
          name: 'studentpost',
          component: StudentPost
        },
        {
          path: 'getpost',
          name: 'studentgetpost',
          component: StudentGetPost
        },
        {
          path: 'putpost',
          name: 'studentputpost',
          component: StudentPutPost
        },
        {
          path: 'delepost',
          name: 'studentdelepost',
          component: StudentDelePost
        },
        {
          path: 'reportpost',
          name: 'studentreportpost',
          component: StudentReportPost
        },
        {
          path: 'getreport',
          name: 'studentgetreport',
          component: StudentGetReport
        },
      ]
    },
    {
      path: '/board/teacher',
      name: 'TeacherBoard',
      component: TeacherBoard,
      children: [
        {
          path: 'getreport',
          name: 'teachergetreport',
          component: AdminGetReport
        },
        {
          path: 'proreport',
          name: 'teacherproreport',
          component: AdminProReport
        },
      ]
    },
  ]
})

export default router
