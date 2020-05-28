import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { DefaultComponent } from './default.component';
import { RouterModule } from '@angular/router';
import { SharedModule } from 'src/app/shared/shared.module';

import { FlexLayoutModule } from '@angular/flex-layout';
import { MatSidenavModule } from '@angular/material/sidenav';
import { MatDividerModule } from '@angular/material/divider';
import { ModulesModule } from 'src/app/modules/modules.module';

@NgModule({
  declarations: [
    DefaultComponent,

  ],
  imports: [
    CommonModule,
    RouterModule,
    SharedModule,
    ModulesModule,
    MatSidenavModule,
    MatDividerModule,
    FlexLayoutModule,
  ]
})
export class DefaultModule { }
