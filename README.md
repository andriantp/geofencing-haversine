# geofencing-haversine


A simple Go implementation for determining whether a location falls within a specified school zoning radius using geographic coordinates (latitude and longitude).

## Use Case

Determine whether a student's home is within a school's zoning radius.

Example:

```text
School: Example High School
Radius: 2 km

Student Home
      ↓
Calculate Distance
      ↓
In Zone / Out of Zone
```

## Project Structure

```text
.
├── go.mod
├── main.go
├── README.md
└── zonasi
    └── zonasi.go
```

## How It Works

Input:

```text
School Latitude
School Longitude

Student Latitude
Student Longitude

Zone Radius
```

Process:

```text
Calculate distance using the Haversine formula
        ↓
Compare with the zoning radius
        ↓
Determine zoning status
```

Condition:

```text
distance <= radius
    => In Zone

distance > radius
    => Out of Zone
```

## Running the Program

Example:

```bash
go run . -6.123456 106.123456
```

Arguments:

```text
go run . <student_latitude> <student_longitude>
```

Example output:

```text
Distance: 1.23 km, In Zone: true
```

## Configure School Location

The school location is defined in `main.go`.

Example:

```go
schoolLocation := zonasi.Location{
    Latitude:  -6.100000,
    Longitude: 106.100000,
}
```

## Configure Zone Radius

The zoning radius is specified when creating the repository.

Example with a 2 km radius:

```go
repo := zonasi.NewRepository(2)
```

## Using the Zonasi Package

Create a repository:

```go
repo := zonasi.NewRepository(2)
```

Define locations:

```go
school := zonasi.Location{
    Latitude:  -6.100000,
    Longitude: 106.100000,
}

user := zonasi.Location{
    Latitude:  -6.110000,
    Longitude: 106.110000,
}
```

Check zoning status:

```go
result := repo.IsInZone(
    school,
    user,
)
```

Result structure:

```go
type Result struct {
    Distance float64
    InZone   bool
}
```

## Technical Notes

This implementation uses the Haversine formula to calculate the geodesic distance between two geographic coordinates.

The calculated distance represents the straight-line distance over the Earth's surface and does not consider roads, rivers, buildings, or other physical obstacles.
