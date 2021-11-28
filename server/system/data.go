package system

import (
	"database/sql"
)

type ListBalancersResult struct {
	Id             int64 `json:"id"`
	Total_machines int64 `json:"total_machines"`
	Machines_id    int64 `json:"machines_id"`
}

type ListBalancersResponse struct {
	Id                 int64   `json:"id"`
	UsedMachines       []int64 `json:"usedMachines"`
	TotalMachinesCount int64   `json:"totalMachinesCount"`
}

type Balancer struct {
	Id  int64  `json:"id"`
	Job string `json:"job"`
}

func balancerInSlice(a int64, list []*ListBalancersResponse) int {
	for i, b := range list {
		if b.Id == a {
			return i
		}
	}
	return -1
}

type Store struct {
	Db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{Db: db}
}

func (s *Store) ListBalancers() ([]*ListBalancersResponse, error) {
	rows, err := s.Db.Query("SELECT balancers.id, ( SELECT COUNT(*) FROM machines WHERE machines.balancer_id = balancers.id ) as total_machines, machines.id as machine_id FROM balancers INNER JOIN machines ON machines.balancer_id = balancers.id WHERE machines.working = TRUE GROUP BY balancers.id, machines.id LIMIT 200")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var res []*ListBalancersResponse
	for rows.Next() {
		var c ListBalancersResult
		if err := rows.Scan(&c.Id, &c.Total_machines, &c.Machines_id); err != nil {
			return nil, err
		}
		ind := balancerInSlice(c.Id, res)
		if ind != -1 {
			res[ind].UsedMachines = append(res[ind].UsedMachines, c.Machines_id)
		} else {
			balancer := &ListBalancersResponse{c.Id, make([]int64, c.Machines_id), c.Total_machines}
			res = append(res, balancer)
		}
	}
	if res == nil {
		res = make([]*ListBalancersResponse, 0)
	}
	return res, nil
}

type Machine struct {
	Id          int64 `json:"id"`
	Working     bool  `json:"working"`
	Balancer_id int64 `json:"id"`
}

func (s *Store) UpdateMachineStatus(id int64, working bool) error {
	_, err := s.Db.Exec("UPDATE machines SET working = $1 where id = $2", working, id)
	return err
}
