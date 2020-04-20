import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import {CallbackComponent} from "./components/callback/callback.component";
import {LoginComponent} from "./components/login/login.component";
import {DashboardComponent} from "./components/dashboard/dashboard.component";
import {DeployComponent} from "./components/deploy/deploy.component";
import {DataConfigurationComponent} from "./components/data-configuration/data-configuration.component";


const routes: Routes = [
  { path: 'callback', component: CallbackComponent },
  { path: 'login', component: LoginComponent },
  { path: 'dashboard',
    component: DashboardComponent, // this is the component with the <router-outlet> in the template
    children: [
      {
        path: 'deploy', // child route path
        component: DeployComponent
      },
      {
        path: 'data-configuration',
        component: DataConfigurationComponent
      }
    ] }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
