import {Component} from '@angular/core';

@Component({
  selector: 'app-root',
  template: `
<section class="section">
    <router-outlet></router-outlet>
</section>
  `,
  styles: []
})
export class AppComponent {
  title = 'CV';
}
