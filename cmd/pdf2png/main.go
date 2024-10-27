// Copyright (c) 2024 Michael D Henderson. All rights reserved.

// Package main implements a command-line tool to extract images from a PDF file and save them as individual PNG files.
//
// The PDF documents are not included in the repository because they are the property of the creators, Crab Fragment Labs.
// You must fetch the documents from the Crab Fragment Labs website (https://crabfragmentlabs.com/stage-blood).
package main

import (
	"log"
	"path/filepath"

	"github.com/gen2brain/go-fitz"
)

func main() {
	log.SetFlags(log.Lshortfile)

	pdfPath := "../cfl"          // path to the downloaded PDF documents
	outputDir := "../cfl/images" // path to save the extracted images

	log.Printf("extractingactor cards\n")
	if doc, err := fitz.New(filepath.Join(pdfPath, "ActorDeckPnP-v1.4.pdf")); err != nil {
		log.Fatalf("error: opening PDF: %v\n", err)
	} else if err := extractActorCardsFromPDF(doc, filepath.Join(outputDir, "actors")); err != nil {
		log.Fatalf("error extracting actor cards: %v\n", err)
	} else {
		_ = doc.Close()
	}

	log.Printf("extractingplay cards\n")
	if doc, err := fitz.New(filepath.Join(pdfPath, "PlayDeckPnP-v1.4.pdf")); err != nil {
		log.Fatalf("error: opening PDF: %v\n", err)
	} else if err := extractPlayCardsFromPDF(doc, filepath.Join(outputDir, "plays")); err != nil {
		log.Fatalf("error extracting play cards: %v\n", err)
	} else {
		_ = doc.Close()
	}

	log.Printf("extractingfavor tokens\n")
	if doc, err := fitz.New(filepath.Join(pdfPath, "StageBloodTokens-v1.4.pdf")); err != nil {
		log.Fatalf("error: opening PDF: %v\n", err)
	} else if err := extractFavorTokensFromPDF(doc, filepath.Join(outputDir, "tokens")); err != nil {
		log.Fatalf("error extracting favor tokens: %v\n", err)
	} else {
		_ = doc.Close()
	}
}
