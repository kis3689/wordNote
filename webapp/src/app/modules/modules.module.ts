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



@NgModule({
  declarations: [
    WordnoteComponent,
    WordDialogComponent,
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
    ReactiveFormsModule
  ]
})
export class ModulesModule { }
