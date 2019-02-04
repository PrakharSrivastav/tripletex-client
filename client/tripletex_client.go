// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/PrakharSrivastav/Petstore/client/accommodation_allowance"
	"github.com/PrakharSrivastav/Petstore/client/account"
	"github.com/PrakharSrivastav/Petstore/client/accounting_period"
	"github.com/PrakharSrivastav/Petstore/client/activity"
	"github.com/PrakharSrivastav/Petstore/client/address"
	"github.com/PrakharSrivastav/Petstore/client/altinn"
	"github.com/PrakharSrivastav/Petstore/client/annual_account"
	"github.com/PrakharSrivastav/Petstore/client/bank"
	"github.com/PrakharSrivastav/Petstore/client/category"
	"github.com/PrakharSrivastav/Petstore/client/close_group"
	"github.com/PrakharSrivastav/Petstore/client/company"
	"github.com/PrakharSrivastav/Petstore/client/consumer"
	"github.com/PrakharSrivastav/Petstore/client/contact"
	"github.com/PrakharSrivastav/Petstore/client/cost"
	"github.com/PrakharSrivastav/Petstore/client/cost_category"
	"github.com/PrakharSrivastav/Petstore/client/country"
	"github.com/PrakharSrivastav/Petstore/client/currency"
	"github.com/PrakharSrivastav/Petstore/client/customer"
	"github.com/PrakharSrivastav/Petstore/client/department"
	"github.com/PrakharSrivastav/Petstore/client/details"
	"github.com/PrakharSrivastav/Petstore/client/division"
	"github.com/PrakharSrivastav/Petstore/client/document"
	"github.com/PrakharSrivastav/Petstore/client/employee"
	"github.com/PrakharSrivastav/Petstore/client/employment"
	"github.com/PrakharSrivastav/Petstore/client/employment_type"
	"github.com/PrakharSrivastav/Petstore/client/entitlement"
	"github.com/PrakharSrivastav/Petstore/client/entry"
	"github.com/PrakharSrivastav/Petstore/client/event"
	"github.com/PrakharSrivastav/Petstore/client/external"
	"github.com/PrakharSrivastav/Petstore/client/holiday"
	"github.com/PrakharSrivastav/Petstore/client/inventory"
	"github.com/PrakharSrivastav/Petstore/client/invoice"
	"github.com/PrakharSrivastav/Petstore/client/leave_of_absence"
	"github.com/PrakharSrivastav/Petstore/client/leave_of_absence_type"
	"github.com/PrakharSrivastav/Petstore/client/ledger"
	"github.com/PrakharSrivastav/Petstore/client/match"
	"github.com/PrakharSrivastav/Petstore/client/mileage_allowance"
	"github.com/PrakharSrivastav/Petstore/client/municipality"
	"github.com/PrakharSrivastav/Petstore/client/next_of_kin"
	"github.com/PrakharSrivastav/Petstore/client/occupation_code"
	"github.com/PrakharSrivastav/Petstore/client/order"
	"github.com/PrakharSrivastav/Petstore/client/orderline"
	"github.com/PrakharSrivastav/Petstore/client/participant"
	"github.com/PrakharSrivastav/Petstore/client/passenger"
	"github.com/PrakharSrivastav/Petstore/client/payment_type"
	"github.com/PrakharSrivastav/Petstore/client/payment_type_out"
	"github.com/PrakharSrivastav/Petstore/client/payslip"
	"github.com/PrakharSrivastav/Petstore/client/per_diem_compensation"
	"github.com/PrakharSrivastav/Petstore/client/posting"
	"github.com/PrakharSrivastav/Petstore/client/product"
	"github.com/PrakharSrivastav/Petstore/client/project"
	"github.com/PrakharSrivastav/Petstore/client/prospect"
	"github.com/PrakharSrivastav/Petstore/client/rate"
	"github.com/PrakharSrivastav/Petstore/client/rate_category"
	"github.com/PrakharSrivastav/Petstore/client/rate_category_group"
	"github.com/PrakharSrivastav/Petstore/client/reconciliation"
	"github.com/PrakharSrivastav/Petstore/client/reminder"
	"github.com/PrakharSrivastav/Petstore/client/remuneration_type"
	"github.com/PrakharSrivastav/Petstore/client/session"
	"github.com/PrakharSrivastav/Petstore/client/settings"
	"github.com/PrakharSrivastav/Petstore/client/standard_time"
	"github.com/PrakharSrivastav/Petstore/client/statement"
	"github.com/PrakharSrivastav/Petstore/client/subscription"
	"github.com/PrakharSrivastav/Petstore/client/supplier"
	"github.com/PrakharSrivastav/Petstore/client/time_clock"
	"github.com/PrakharSrivastav/Petstore/client/transaction"
	"github.com/PrakharSrivastav/Petstore/client/travel_expense"
	"github.com/PrakharSrivastav/Petstore/client/type_operations"
	"github.com/PrakharSrivastav/Petstore/client/unit"
	"github.com/PrakharSrivastav/Petstore/client/vat_type"
	"github.com/PrakharSrivastav/Petstore/client/voucher"
	"github.com/PrakharSrivastav/Petstore/client/voucher_type"
	"github.com/PrakharSrivastav/Petstore/client/working_hours_scheme"
)

// Default tripletex HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "tripletex.no"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/v2"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http"}

// NewHTTPClient creates a new tripletex HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Tripletex {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new tripletex HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Tripletex {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new tripletex client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Tripletex {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Tripletex)
	cli.Transport = transport

	cli.AccommodationAllowance = accommodation_allowance.New(transport, formats)

	cli.Account = account.New(transport, formats)

	cli.AccountingPeriod = accounting_period.New(transport, formats)

	cli.Activity = activity.New(transport, formats)

	cli.Address = address.New(transport, formats)

	cli.Altinn = altinn.New(transport, formats)

	cli.AnnualAccount = annual_account.New(transport, formats)

	cli.Bank = bank.New(transport, formats)

	cli.Category = category.New(transport, formats)

	cli.CloseGroup = close_group.New(transport, formats)

	cli.Company = company.New(transport, formats)

	cli.Consumer = consumer.New(transport, formats)

	cli.Contact = contact.New(transport, formats)

	cli.Cost = cost.New(transport, formats)

	cli.CostCategory = cost_category.New(transport, formats)

	cli.Country = country.New(transport, formats)

	cli.Currency = currency.New(transport, formats)

	cli.Customer = customer.New(transport, formats)

	cli.Department = department.New(transport, formats)

	cli.Details = details.New(transport, formats)

	cli.Division = division.New(transport, formats)

	cli.Document = document.New(transport, formats)

	cli.Employee = employee.New(transport, formats)

	cli.Employment = employment.New(transport, formats)

	cli.EmploymentType = employment_type.New(transport, formats)

	cli.Entitlement = entitlement.New(transport, formats)

	cli.Entry = entry.New(transport, formats)

	cli.Event = event.New(transport, formats)

	cli.External = external.New(transport, formats)

	cli.Holiday = holiday.New(transport, formats)

	cli.Inventory = inventory.New(transport, formats)

	cli.Invoice = invoice.New(transport, formats)

	cli.LeaveOfAbsence = leave_of_absence.New(transport, formats)

	cli.LeaveOfAbsenceType = leave_of_absence_type.New(transport, formats)

	cli.Ledger = ledger.New(transport, formats)

	cli.Match = match.New(transport, formats)

	cli.MileageAllowance = mileage_allowance.New(transport, formats)

	cli.Municipality = municipality.New(transport, formats)

	cli.NextOfKin = next_of_kin.New(transport, formats)

	cli.OccupationCode = occupation_code.New(transport, formats)

	cli.Order = order.New(transport, formats)

	cli.Orderline = orderline.New(transport, formats)

	cli.Participant = participant.New(transport, formats)

	cli.Passenger = passenger.New(transport, formats)

	cli.PaymentType = payment_type.New(transport, formats)

	cli.PaymentTypeOut = payment_type_out.New(transport, formats)

	cli.Payslip = payslip.New(transport, formats)

	cli.PerDiemCompensation = per_diem_compensation.New(transport, formats)

	cli.Posting = posting.New(transport, formats)

	cli.Product = product.New(transport, formats)

	cli.Project = project.New(transport, formats)

	cli.Prospect = prospect.New(transport, formats)

	cli.Rate = rate.New(transport, formats)

	cli.RateCategory = rate_category.New(transport, formats)

	cli.RateCategoryGroup = rate_category_group.New(transport, formats)

	cli.Reconciliation = reconciliation.New(transport, formats)

	cli.Reminder = reminder.New(transport, formats)

	cli.RemunerationType = remuneration_type.New(transport, formats)

	cli.Session = session.New(transport, formats)

	cli.Settings = settings.New(transport, formats)

	cli.StandardTime = standard_time.New(transport, formats)

	cli.Statement = statement.New(transport, formats)

	cli.Subscription = subscription.New(transport, formats)

	cli.Supplier = supplier.New(transport, formats)

	cli.TimeClock = time_clock.New(transport, formats)

	cli.Transaction = transaction.New(transport, formats)

	cli.TravelExpense = travel_expense.New(transport, formats)

	cli.TypeOperations = type_operations.New(transport, formats)

	cli.Unit = unit.New(transport, formats)

	cli.VatType = vat_type.New(transport, formats)

	cli.Voucher = voucher.New(transport, formats)

	cli.VoucherType = voucher_type.New(transport, formats)

	cli.WorkingHoursScheme = working_hours_scheme.New(transport, formats)

	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// Tripletex is a client for tripletex
type Tripletex struct {
	AccommodationAllowance *accommodation_allowance.Client

	Account *account.Client

	AccountingPeriod *accounting_period.Client

	Activity *activity.Client

	Address *address.Client

	Altinn *altinn.Client

	AnnualAccount *annual_account.Client

	Bank *bank.Client

	Category *category.Client

	CloseGroup *close_group.Client

	Company *company.Client

	Consumer *consumer.Client

	Contact *contact.Client

	Cost *cost.Client

	CostCategory *cost_category.Client

	Country *country.Client

	Currency *currency.Client

	Customer *customer.Client

	Department *department.Client

	Details *details.Client

	Division *division.Client

	Document *document.Client

	Employee *employee.Client

	Employment *employment.Client

	EmploymentType *employment_type.Client

	Entitlement *entitlement.Client

	Entry *entry.Client

	Event *event.Client

	External *external.Client

	Holiday *holiday.Client

	Inventory *inventory.Client

	Invoice *invoice.Client

	LeaveOfAbsence *leave_of_absence.Client

	LeaveOfAbsenceType *leave_of_absence_type.Client

	Ledger *ledger.Client

	Match *match.Client

	MileageAllowance *mileage_allowance.Client

	Municipality *municipality.Client

	NextOfKin *next_of_kin.Client

	OccupationCode *occupation_code.Client

	Order *order.Client

	Orderline *orderline.Client

	Participant *participant.Client

	Passenger *passenger.Client

	PaymentType *payment_type.Client

	PaymentTypeOut *payment_type_out.Client

	Payslip *payslip.Client

	PerDiemCompensation *per_diem_compensation.Client

	Posting *posting.Client

	Product *product.Client

	Project *project.Client

	Prospect *prospect.Client

	Rate *rate.Client

	RateCategory *rate_category.Client

	RateCategoryGroup *rate_category_group.Client

	Reconciliation *reconciliation.Client

	Reminder *reminder.Client

	RemunerationType *remuneration_type.Client

	Session *session.Client

	Settings *settings.Client

	StandardTime *standard_time.Client

	Statement *statement.Client

	Subscription *subscription.Client

	Supplier *supplier.Client

	TimeClock *time_clock.Client

	Transaction *transaction.Client

	TravelExpense *travel_expense.Client

	TypeOperations *type_operations.Client

	Unit *unit.Client

	VatType *vat_type.Client

	Voucher *voucher.Client

	VoucherType *voucher_type.Client

	WorkingHoursScheme *working_hours_scheme.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Tripletex) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.AccommodationAllowance.SetTransport(transport)

	c.Account.SetTransport(transport)

	c.AccountingPeriod.SetTransport(transport)

	c.Activity.SetTransport(transport)

	c.Address.SetTransport(transport)

	c.Altinn.SetTransport(transport)

	c.AnnualAccount.SetTransport(transport)

	c.Bank.SetTransport(transport)

	c.Category.SetTransport(transport)

	c.CloseGroup.SetTransport(transport)

	c.Company.SetTransport(transport)

	c.Consumer.SetTransport(transport)

	c.Contact.SetTransport(transport)

	c.Cost.SetTransport(transport)

	c.CostCategory.SetTransport(transport)

	c.Country.SetTransport(transport)

	c.Currency.SetTransport(transport)

	c.Customer.SetTransport(transport)

	c.Department.SetTransport(transport)

	c.Details.SetTransport(transport)

	c.Division.SetTransport(transport)

	c.Document.SetTransport(transport)

	c.Employee.SetTransport(transport)

	c.Employment.SetTransport(transport)

	c.EmploymentType.SetTransport(transport)

	c.Entitlement.SetTransport(transport)

	c.Entry.SetTransport(transport)

	c.Event.SetTransport(transport)

	c.External.SetTransport(transport)

	c.Holiday.SetTransport(transport)

	c.Inventory.SetTransport(transport)

	c.Invoice.SetTransport(transport)

	c.LeaveOfAbsence.SetTransport(transport)

	c.LeaveOfAbsenceType.SetTransport(transport)

	c.Ledger.SetTransport(transport)

	c.Match.SetTransport(transport)

	c.MileageAllowance.SetTransport(transport)

	c.Municipality.SetTransport(transport)

	c.NextOfKin.SetTransport(transport)

	c.OccupationCode.SetTransport(transport)

	c.Order.SetTransport(transport)

	c.Orderline.SetTransport(transport)

	c.Participant.SetTransport(transport)

	c.Passenger.SetTransport(transport)

	c.PaymentType.SetTransport(transport)

	c.PaymentTypeOut.SetTransport(transport)

	c.Payslip.SetTransport(transport)

	c.PerDiemCompensation.SetTransport(transport)

	c.Posting.SetTransport(transport)

	c.Product.SetTransport(transport)

	c.Project.SetTransport(transport)

	c.Prospect.SetTransport(transport)

	c.Rate.SetTransport(transport)

	c.RateCategory.SetTransport(transport)

	c.RateCategoryGroup.SetTransport(transport)

	c.Reconciliation.SetTransport(transport)

	c.Reminder.SetTransport(transport)

	c.RemunerationType.SetTransport(transport)

	c.Session.SetTransport(transport)

	c.Settings.SetTransport(transport)

	c.StandardTime.SetTransport(transport)

	c.Statement.SetTransport(transport)

	c.Subscription.SetTransport(transport)

	c.Supplier.SetTransport(transport)

	c.TimeClock.SetTransport(transport)

	c.Transaction.SetTransport(transport)

	c.TravelExpense.SetTransport(transport)

	c.TypeOperations.SetTransport(transport)

	c.Unit.SetTransport(transport)

	c.VatType.SetTransport(transport)

	c.Voucher.SetTransport(transport)

	c.VoucherType.SetTransport(transport)

	c.WorkingHoursScheme.SetTransport(transport)

}
