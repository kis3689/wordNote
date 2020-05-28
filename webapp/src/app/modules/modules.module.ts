import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { WordnoteComponent } from './wordnote/wordnote.component';

import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import {MatTableModule} from "@angular/material/table";
import {MatButtonModule} from "@angular/material/button";
import {MatDialogModule} from "@angular/material/dialog";
import {MatFormFieldModule} from "@angular/material/form-field";
import {FormsModule, ReactiveFormsModule} from "@angular/forms";
import {MatInputModule} from "@angular/material/input";
import {HttpClientModule} from "@angular/common/http";
import {MatToolbarModule} from "@angular/material/toolbar";
import { WordDialogComponent } from './wordnote/word-dialog/word-dialog.component';
import { WordnoteService } from './wordnote/wordnote.service';
import { UsersComponent } from './users/users.component';
import { MatSlideToggleModule } from '@angular/material/slide-toggle'
import { MatSelectModule } from '@angular/material/select'
import { MatDividerModule } from '@angular/material/divider'

@NgModule({
  declarations: [
    WordnoteComponent,
    WordDialogComponent,
    UsersComponent,
  ],
  imports: [
    CommonModule,
    BrowserAnimationsModule,
    MatInputModule,
    MatDialogModule,
    MatTableModule,
    MatButtonModule,
    MatFormFieldModule,
    FormsModule,
    HttpClientModule,
    MatToolbarModule,
    ReactiveFormsModule,
    MatSlideToggleModule,
    MatSelectModule,
    MatDividerModule

  ],
  providers: [
    WordnoteService
  ],
  entryComponents: [
    WordDialogComponent
  ]
})
export class ModulesModule { }
