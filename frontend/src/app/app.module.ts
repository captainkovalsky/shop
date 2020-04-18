import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { FbAuthComponent } from './components/fb-auth/fb-auth.component';
import { CallbackComponent } from './components/callback/callback.component';
import {HTTP_INTERCEPTORS, HttpClientModule} from "@angular/common/http";
import {TokenInterceptor} from "./services/interceptors/token.interceptor";

@NgModule({
  declarations: [
    AppComponent,
    FbAuthComponent,
    CallbackComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
  ],
    providers: [{
      provide: HTTP_INTERCEPTORS,
      useClass: TokenInterceptor,
      multi: true
    }],
  bootstrap: [AppComponent]
})
export class AppModule { }
