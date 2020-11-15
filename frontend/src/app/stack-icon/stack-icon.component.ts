import {Component, Input, OnInit} from '@angular/core';

@Component({
  selector: 'app-stack-icon',
  template: `
    <p>
      stack-icon works!
    </p>
  `,
  styles: [
  ]
})
export class StackIconComponent implements OnInit {
  @Input() stack: string[] = []
  constructor() { }

  ngOnInit(): void {
  }

}
