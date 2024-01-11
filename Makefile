start-backend:
	docker compose up --build

remove:
	docker rm hahaton-gateway-1
	docker rm hahaton-bell-server-1
	docker rm hahaton-audio-storage-1
	docker rm hahaton-trans-service-1
	docker rm hahaton-postgres-1
	docker rm hahaton-whisper-1

stop:
	docker stop hahaton-gateway-1
	docker stop hahaton-bell-server-1
	docker stop hahaton-audio-storage-1
	docker stop hahaton-trans-service-1
	docker stop hahaton-postgres-1
	docker stop hahaton-whisper-1


#npm install hahaton/frontend
#npm run dev -- --open