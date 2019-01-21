#!/bin/bash

echo "Creating project"
#curl -X POST http://localhost:8080/v1beta1/projects -H "Content-Type: application/json" --data '{"name":"projects/image-signing"}'
echo ""

echo "Creating note"
# @TODO Adjust path to @note.json
#curl -X POST "http://localhost:8080/v1beta1/projects/image-signing/notes" -d @note.json
echo ""

echo "Creating occurrence"
# @TODO Adjust path to occurrence.json
#curl -X POST 'http://127.0.0.1:8080/v1beta1/projects/image-signing/occurrences' -H "Content-Type: application/json" -d @occurrence.json
echo ""

echo "Starting image-signature-webhook"
go run ../main.go -tls-cert ../../pki/image-signature-webhook.pem -tls-key ../../pki/image-signature-webhook-key.pem -grafeas http://localhost:8080