## Library Management System
This is a simple console-based library management system implemented in Go.

### Core Features
```
### Simultaneous Requests
* Add a book
* Remove a book
* Borrow a book
* Return a borrowed book
* Reserve a book

```
## Concurrent Book Reservation Feature
Handling Multiple Reservation Requests concurrently with gorountines, mutexes, channels, time and wait groups. 

### How I handled Simultaneous Requests
```
### Simultaneous Requests
* Achieved using **Goroutines** and **Channels**
* Processes multiple incoming `ReservationRequest` tasks concurrently
* Uses dedicated workers
```

### How I Ensured Data Safety
```
### Data Safety
- Uses `sync.Mutex` to prevent race conditions
- Shared state: `Books`, `Members`, `Reservations`
- Lock must be acquired before modification
```

## Core Methods

```
## Core Method
### `ReserveBook(bookID int, memberID int) error`
### `RemoveBook(bookID int, memberID int) error`

```

#### Logic Flow
```
1. Acquire lock using `library.Mutex.Lock()`
2. Check book status: return error if `Borrowed`
3. Check for existing reservations:
   - Return error if unexpired
   - Remove the reservation if expired
4. Update book status to `Reserved`
5. Set up Auto-Cancellation Timer with `time.AfterFunc`
6. Launch asynchronous function to borrow book


```

#### Running Program
```
### Use the command below to run the program
go run main.go